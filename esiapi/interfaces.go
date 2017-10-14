package esiapi

import (
	"github.com/antihax/goesi/esi"
	"net/http"
)

type AllianceApi interface {
	GetAlliancesAllianceId(allianceId int32, localVarOptionals map[string]interface{}) (esi.GetAlliancesAllianceIdOk, *http.Response, error)
}

type CorporationApi interface {
	GetCorporationsCorporationId(corporationId int32, localVarOptionals map[string]interface{}) (esi.GetCorporationsCorporationIdOk, *http.Response, error)
}

type CharacterApi interface {
	GetCharactersCharacterId(characterId int32, localVarOptionals map[string]interface{}) (esi.GetCharactersCharacterIdOk, *http.Response, error)
}
