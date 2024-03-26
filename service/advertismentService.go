package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"rest/dto"
	"rest/errcstm"
	"rest/jwt"
	"rest/model"
	"rest/queries"
	"sort"
	"time"
)

type AdvService struct {
	db *sql.DB
}

func NewAdvService(db *sql.DB) *AdvService {
	return &AdvService{
		db: db,
	}
}

func (advs *AdvService) AddAdv(r *http.Request, userService *UserService) (dto.AdvertisementDto, int, error) {
	var adv model.Advertisement

	if r.Method != "POST" {
		return adv.ModelToDto(), http.StatusBadRequest, errors.New(errcstm.REQUEST_TYPE)
	}

	err := json.NewDecoder(r.Body).Decode(&adv)
	if err != nil {
		return adv.ModelToDto(), http.StatusInternalServerError, errors.New(errcstm.REQUEST_BODY + err.Error())
	}

	if err = adv.FilterFields(); err != nil {
		return adv.ModelToDto(), http.StatusInternalServerError, errors.New(errcstm.REQUEST_BODY + err.Error())
	}

	adv.CreateId()
	adv.CreatedDtTm = time.Now()

	adv.UserId, err = jwt.VerifyJWT(r)
	if err != nil {
		return adv.ModelToDto(), http.StatusForbidden, errors.New(errcstm.USER_SELECT + err.Error())
	}

	_, err = advs.db.Exec(queries.INSERT_ADV, adv.ID, adv.CreatedDtTm, adv.UserId, adv.Title, adv.Content, adv.ImgPth, adv.Price)
	if err != nil {
		return adv.ModelToDto(), http.StatusInternalServerError, errors.New(errcstm.ADVERTISMENT_INSERT + err.Error())
	}
	return adv.ModelToDto(), http.StatusOK, nil
}

func (advs *AdvService) GetAdv(r *http.Request, userService *UserService) ([]dto.AdvertisementSendDto, int, error) {
	advertisements := []model.AdvertisementGetDto{}
	advsDto := []dto.AdvertisementSendDto{}

	if r.Method != "GET" {
		return advsDto, http.StatusBadRequest, errors.New(errcstm.REQUEST_TYPE)
	}

	rows, err := advs.db.Query(queries.SELECT_ADV)
	if err != nil {
		return advsDto, http.StatusInternalServerError, errors.New(errcstm.ADVERTISMENT_INSERT + err.Error())
	}

	for rows.Next() {
		var advGet model.AdvertisementGetDto
		err := rows.Scan(&advGet.UserLogin, &advGet.Adv.ID, &advGet.Adv.CreatedDtTm, &advGet.Adv.UserId, &advGet.Adv.Title, &advGet.Adv.Content, &advGet.Adv.ImgPth, &advGet.Adv.Price)
		if err != nil {
			return advsDto, http.StatusInternalServerError, errors.New(errcstm.ADVERTISMENT_SELECT + err.Error())
		}
		advertisements = append(advertisements, advGet)
	}

	sort.Slice(advertisements, func(i, j int) bool {
		return advertisements[i].Adv.CreatedDtTm.Before(advertisements[j].Adv.CreatedDtTm)
	})

	userId, _ := jwt.VerifyJWT(r)
	for _, advertisment := range advertisements {
		advsDto = append(advsDto, advertisment.GetToSend(userId))
	}

	return advsDto, http.StatusOK, nil
}
