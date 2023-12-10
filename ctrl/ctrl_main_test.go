package ctrl

import (
	"fmt"
	"github.com/Wilm0rien/omip/model"
	"github.com/Wilm0rien/omip/util"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	// test token generated via http://jwtbuilder.jamiekurtz.com/
	// only two fields are relevant for this test
	// sub  : CHARACTER:EVE:2115636466
	// name : Ion of Chios
	// see GetCharInfo() function for jwt extraction of token
	dummyToken := `eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJvbWlwIHRlc3QgdG9rZW4iLCJpYXQiOjE2MzU3NTE2NjMsImV4cCI6MTY2NzI4NzY2MywiYXVkIjoid3d3LmV2ZW9ubGluZS5jb20iLCJzdWIiOiJDSEFSQUNURVI6RVZFOjIxMTU2MzY0NjYiLCJuYW1lIjoiSW9uIG9mIENoaW9zIiwiRW1haWwiOiJqcm9ja2V0QGV4YW1wbGUuY29tIn0.kbAkeoDeGh3Hh5mVtKJNl-vJScbbkOOlTYTs1mR91ZY`

	CtrlTestEnable = true
	model.DeleteDb(model.DbNameCtrlTest)
	modelObj := model.NewModel(model.DbNameCtrlTest, false)
	ctrlObj := NewCtrl(modelObj)
	waitForAuth := make(chan string)
	expiresOn := util.UnixTS2DateTimeStr(time.Now().Add(1199 * time.Second).Unix())

	response := func(req *http.Request) (bodyBytes []byte, err error, resp *http.Response) {
		resp = &http.Response{
			StatusCode: http.StatusOK,
		}
		switch req.URL.String() {
		case "https://login.eveonline.com/v2/oauth/token":
			bodyBytes = []byte("{\"access_token\":\"" + dummyToken + "\",\"expires_in\":1199,\"token_type\":\"Bearer\",\"refresh_token\":\"refresh_token_dummytoken\"}")
		case "https://login.eveonline.com/oauth/verify":
			resultString := fmt.Sprintf("{\"CharacterID\":2115636466,\"CharacterName\":\"Ion of Chios\",\"ExpiresOn\":\"%s\",\"Scopes\":\"publicData esi-wallet.read_character_wallet.v1 esi-wallet.read_corporation_wallet.v1 esi-universe.read_structures.v1 esi-killmails.read_killmails.v1 esi-corporations.read_corporation_membership.v1 esi-corporations.read_structures.v1 esi-industry.read_character_jobs.v1 esi-contracts.read_character_contracts.v1 esi-killmails.read_corporation_killmails.v1 esi-corporations.track_members.v1 esi-wallet.read_corporation_wallets.v1 esi-characters.read_notifications.v1 esi-contracts.read_corporation_contracts.v1 esi-corporations.read_starbases.v1 esi-industry.read_corporation_jobs.v1\",\"TokenType\":\"Character\",\"CharacterOwnerHash\":\"dummyhash=\",\"IntellectualProperty\":\"EVE\"}",
				expiresOn)
			bodyBytes = []byte(resultString)
		case "https://esi.evetech.net/v5/characters/2115636466":
			bodyBytes = []byte("{\"ancestry_id\":9,\"birthday\":\"2019-08-28T18:48:02Z\",\"bloodline_id\":2,\"corporation_id\":98627127,\"description\":\"\",\"gender\":\"male\",\"name\":\"Ion of Chios\",\"race_id\":1,\"security_status\":0.0}")
		case "https://esi.evetech.net/v2/corporations/98627127/roles/?datasource=tranquility":
			bodyBytes = []byte("[{\"character_id\":95067057,\"grantable_roles\":[\"Director\",\"Personnel_Manager\",\"Accountant\",\"Security_Officer\",\"Factory_Manager\",\"Station_Manager\",\"Auditor\",\"Hangar_Take_1\",\"Hangar_Take_2\",\"Hangar_Take_3\",\"Hangar_Take_4\",\"Hangar_Take_5\",\"Hangar_Take_6\",\"Hangar_Take_7\",\"Hangar_Query_1\",\"Hangar_Query_2\",\"Hangar_Query_3\",\"Hangar_Query_4\",\"Hangar_Query_5\",\"Hangar_Query_6\",\"Hangar_Query_7\",\"Account_Take_1\",\"Account_Take_2\",\"Account_Take_3\",\"Account_Take_4\",\"Account_Take_5\",\"Account_Take_6\",\"Account_Take_7\",\"Diplomat\",\"Config_Equipment\",\"Container_Take_1\",\"Container_Take_2\",\"Container_Take_3\",\"Container_Take_4\",\"Container_Take_5\",\"Container_Take_6\",\"Container_Take_7\",\"Rent_Office\",\"Rent_Factory_Facility\",\"Rent_Research_Facility\",\"Junior_Accountant\",\"Config_Starbase_Equipment\",\"Trader\",\"Communications_Officer\",\"Contract_Manager\",\"Starbase_Defense_Operator\",\"Starbase_Fuel_Technician\",\"Fitting_Manager\"],\"grantable_roles_at_base\":[\"Director\",\"Personnel_Manager\",\"Accountant\",\"Security_Officer\",\"Factory_Manager\",\"Station_Manager\",\"Auditor\",\"Hangar_Take_1\",\"Hangar_Take_2\",\"Hangar_Take_3\",\"Hangar_Take_4\",\"Hangar_Take_5\",\"Hangar_Take_6\",\"Hangar_Take_7\",\"Hangar_Query_1\",\"Hangar_Query_2\",\"Hangar_Query_3\",\"Hangar_Query_4\",\"Hangar_Query_5\",\"Hangar_Query_6\",\"Hangar_Query_7\",\"Account_Take_1\",\"Account_Take_2\",\"Account_Take_3\",\"Account_Take_4\",\"Account_Take_5\",\"Account_Take_6\",\"Account_Take_7\",\"Diplomat\",\"Config_Equipment\",\"Container_Take_1\",\"Container_Take_2\",\"Container_Take_3\",\"Container_Take_4\",\"Container_Take_5\",\"Container_Take_6\",\"Container_Take_7\",\"Rent_Office\",\"Rent_Factory_Facility\",\"Rent_Research_Facility\",\"Junior_Accountant\",\"Config_Starbase_Equipment\",\"Trader\",\"Communications_Officer\",\"Contract_Manager\",\"Starbase_Defense_Operator\",\"Starbase_Fuel_Technician\",\"Fitting_Manager\"],\"grantable_roles_at_hq\":[\"Director\",\"Personnel_Manager\",\"Accountant\",\"Security_Officer\",\"Factory_Manager\",\"Station_Manager\",\"Auditor\",\"Hangar_Take_1\",\"Hangar_Take_2\",\"Hangar_Take_3\",\"Hangar_Take_4\",\"Hangar_Take_5\",\"Hangar_Take_6\",\"Hangar_Take_7\",\"Hangar_Query_1\",\"Hangar_Query_2\",\"Hangar_Query_3\",\"Hangar_Query_4\",\"Hangar_Query_5\",\"Hangar_Query_6\",\"Hangar_Query_7\",\"Account_Take_1\",\"Account_Take_2\",\"Account_Take_3\",\"Account_Take_4\",\"Account_Take_5\",\"Account_Take_6\",\"Account_Take_7\",\"Diplomat\",\"Config_Equipment\",\"Container_Take_1\",\"Container_Take_2\",\"Container_Take_3\",\"Container_Take_4\",\"Container_Take_5\",\"Container_Take_6\",\"Container_Take_7\",\"Rent_Office\",\"Rent_Factory_Facility\",\"Rent_Research_Facility\",\"Junior_Accountant\",\"Config_Starbase_Equipment\",\"Trader\",\"Communications_Officer\",\"Contract_Manager\",\"Starbase_Defense_Operator\",\"Starbase_Fuel_Technician\",\"Fitting_Manager\"],\"grantable_roles_at_other\":[\"Director\",\"Personnel_Manager\",\"Accountant\",\"Security_Officer\",\"Factory_Manager\",\"Station_Manager\",\"Auditor\",\"Hangar_Take_1\",\"Hangar_Take_2\",\"Hangar_Take_3\",\"Hangar_Take_4\",\"Hangar_Take_5\",\"Hangar_Take_6\",\"Hangar_Take_7\",\"Hangar_Query_1\",\"Hangar_Query_2\",\"Hangar_Query_3\",\"Hangar_Query_4\",\"Hangar_Query_5\",\"Hangar_Query_6\",\"Hangar_Query_7\",\"Account_Take_1\",\"Account_Take_2\",\"Account_Take_3\",\"Account_Take_4\",\"Account_Take_5\",\"Account_Take_6\",\"Account_Take_7\",\"Diplomat\",\"Config_Equipment\",\"Container_Take_1\",\"Container_Take_2\",\"Container_Take_3\",\"Container_Take_4\",\"Container_Take_5\",\"Container_Take_6\",\"Container_Take_7\",\"Rent_Office\",\"Rent_Factory_Facility\",\"Rent_Research_Facility\",\"Junior_Accountant\",\"Config_Starbase_Equipment\",\"Trader\",\"Communications_Officer\",\"Contract_Manager\",\"Starbase_Defense_Operator\",\"Starbase_Fuel_Technician\",\"Fitting_Manager\"],\"roles\":[\"Director\",\"Personnel_Manager\",\"Accountant\",\"Security_Officer\",\"Factory_Manager\",\"Station_Manager\",\"Auditor\",\"Hangar_Take_1\",\"Hangar_Take_2\",\"Hangar_Take_3\",\"Hangar_Take_4\",\"Hangar_Take_5\",\"Hangar_Take_6\",\"Hangar_Take_7\",\"Hangar_Query_1\",\"Hangar_Query_2\",\"Hangar_Query_3\",\"Hangar_Query_4\",\"Hangar_Query_5\",\"Hangar_Query_6\",\"Hangar_Query_7\",\"Account_Take_1\",\"Account_Take_2\",\"Account_Take_3\",\"Account_Take_4\",\"Account_Take_5\",\"Account_Take_6\",\"Account_Take_7\",\"Diplomat\",\"Config_Equipment\",\"Container_Take_1\",\"Container_Take_2\",\"Container_Take_3\",\"Container_Take_4\",\"Container_Take_5\",\"Container_Take_6\",\"Container_Take_7\",\"Rent_Office\",\"Rent_Factory_Facility\",\"Rent_Research_Facility\",\"Junior_Accountant\",\"Config_Starbase_Equipment\",\"Trader\",\"Communications_Officer\",\"Contract_Manager\",\"Starbase_Defense_Operator\",\"Starbase_Fuel_Technician\",\"Fitting_Manager\"],\"roles_at_base\":[\"Director\",\"Personnel_Manager\",\"Accountant\",\"Security_Officer\",\"Factory_Manager\",\"Station_Manager\",\"Auditor\",\"Hangar_Take_1\",\"Hangar_Take_2\",\"Hangar_Take_3\",\"Hangar_Take_4\",\"Hangar_Take_5\",\"Hangar_Take_6\",\"Hangar_Take_7\",\"Hangar_Query_1\",\"Hangar_Query_2\",\"Hangar_Query_3\",\"Hangar_Query_4\",\"Hangar_Query_5\",\"Hangar_Query_6\",\"Hangar_Query_7\",\"Account_Take_1\",\"Account_Take_2\",\"Account_Take_3\",\"Account_Take_4\",\"Account_Take_5\",\"Account_Take_6\",\"Account_Take_7\",\"Diplomat\",\"Config_Equipment\",\"Container_Take_1\",\"Container_Take_2\",\"Container_Take_3\",\"Container_Take_4\",\"Container_Take_5\",\"Container_Take_6\",\"Container_Take_7\",\"Rent_Office\",\"Rent_Factory_Facility\",\"Rent_Research_Facility\",\"Junior_Accountant\",\"Config_Starbase_Equipment\",\"Trader\",\"Communications_Officer\",\"Contract_Manager\",\"Starbase_Defense_Operator\",\"Starbase_Fuel_Technician\",\"Fitting_Manager\"],\"roles_at_hq\":[\"Director\",\"Personnel_Manager\",\"Accountant\",\"Security_Officer\",\"Factory_Manager\",\"Station_Manager\",\"Auditor\",\"Hangar_Take_1\",\"Hangar_Take_2\",\"Hangar_Take_3\",\"Hangar_Take_4\",\"Hangar_Take_5\",\"Hangar_Take_6\",\"Hangar_Take_7\",\"Hangar_Query_1\",\"Hangar_Query_2\",\"Hangar_Query_3\",\"Hangar_Query_4\",\"Hangar_Query_5\",\"Hangar_Query_6\",\"Hangar_Query_7\",\"Account_Take_1\",\"Account_Take_2\",\"Account_Take_3\",\"Account_Take_4\",\"Account_Take_5\",\"Account_Take_6\",\"Account_Take_7\",\"Diplomat\",\"Config_Equipment\",\"Container_Take_1\",\"Container_Take_2\",\"Container_Take_3\",\"Container_Take_4\",\"Container_Take_5\",\"Container_Take_6\",\"Container_Take_7\",\"Rent_Office\",\"Rent_Factory_Facility\",\"Rent_Research_Facility\",\"Junior_Accountant\",\"Config_Starbase_Equipment\",\"Trader\",\"Communications_Officer\",\"Contract_Manager\",\"Starbase_Defense_Operator\",\"Starbase_Fuel_Technician\",\"Fitting_Manager\"],\"roles_at_other\":[\"Director\",\"Personnel_Manager\",\"Accountant\",\"Security_Officer\",\"Factory_Manager\",\"Station_Manager\",\"Auditor\",\"Hangar_Take_1\",\"Hangar_Take_2\",\"Hangar_Take_3\",\"Hangar_Take_4\",\"Hangar_Take_5\",\"Hangar_Take_6\",\"Hangar_Take_7\",\"Hangar_Query_1\",\"Hangar_Query_2\",\"Hangar_Query_3\",\"Hangar_Query_4\",\"Hangar_Query_5\",\"Hangar_Query_6\",\"Hangar_Query_7\",\"Account_Take_1\",\"Account_Take_2\",\"Account_Take_3\",\"Account_Take_4\",\"Account_Take_5\",\"Account_Take_6\",\"Account_Take_7\",\"Diplomat\",\"Config_Equipment\",\"Container_Take_1\",\"Container_Take_2\",\"Container_Take_3\",\"Container_Take_4\",\"Container_Take_5\",\"Container_Take_6\",\"Container_Take_7\",\"Rent_Office\",\"Rent_Factory_Facility\",\"Rent_Research_Facility\",\"Junior_Accountant\",\"Config_Starbase_Equipment\",\"Trader\",\"Communications_Officer\",\"Contract_Manager\",\"Starbase_Defense_Operator\",\"Starbase_Fuel_Technician\",\"Fitting_Manager\"]},{\"character_id\":95281762,\"grantable_roles\":[],\"grantable_roles_at_base\":[],\"grantable_roles_at_hq\":[],\"grantable_roles_at_other\":[],\"roles\":[],\"roles_at_base\":[],\"roles_at_hq\":[],\"roles_at_other\":[]},{\"character_id\":2113199519,\"grantable_roles\":[],\"grantable_roles_at_base\":[],\"grantable_roles_at_hq\":[],\"grantable_roles_at_other\":[],\"roles\":[],\"roles_at_base\":[],\"roles_at_hq\":[],\"roles_at_other\":[]},{\"character_id\":2114367476,\"grantable_roles\":[],\"grantable_roles_at_base\":[],\"grantable_roles_at_hq\":[],\"grantable_roles_at_other\":[],\"roles\":[],\"roles_at_base\":[],\"roles_at_hq\":[],\"roles_at_other\":[]},{\"character_id\":2114908444,\"grantable_roles\":[],\"grantable_roles_at_base\":[],\"grantable_roles_at_hq\":[],\"grantable_roles_at_other\":[],\"roles\":[],\"roles_at_base\":[],\"roles_at_hq\":[],\"roles_at_other\":[]},{\"character_id\":2115417359,\"grantable_roles\":[],\"grantable_roles_at_base\":[],\"grantable_roles_at_hq\":[],\"grantable_roles_at_other\":[],\"roles\":[],\"roles_at_base\":[],\"roles_at_hq\":[],\"roles_at_other\":[]},{\"character_id\":2115448095,\"grantable_roles\":[],\"grantable_roles_at_base\":[],\"grantable_roles_at_hq\":[],\"grantable_roles_at_other\":[],\"roles\":[],\"roles_at_base\":[],\"roles_at_hq\":[],\"roles_at_other\":[]},{\"character_id\":2115636466,\"grantable_roles\":[],\"grantable_roles_at_base\":[],\"grantable_roles_at_hq\":[],\"grantable_roles_at_other\":[],\"roles\":[\"Director\"],\"roles_at_base\":[],\"roles_at_hq\":[],\"roles_at_other\":[]},{\"character_id\":2115692519,\"grantable_roles\":[],\"grantable_roles_at_base\":[],\"grantable_roles_at_hq\":[],\"grantable_roles_at_other\":[],\"roles\":[],\"roles_at_base\":[],\"roles_at_hq\":[],\"roles_at_other\":[]},{\"character_id\":2115692575,\"grantable_roles\":[],\"grantable_roles_at_base\":[],\"grantable_roles_at_hq\":[],\"grantable_roles_at_other\":[],\"roles\":[],\"roles_at_base\":[],\"roles_at_hq\":[],\"roles_at_other\":[]},{\"character_id\":2115714045,\"grantable_roles\":[],\"grantable_roles_at_base\":[],\"grantable_roles_at_hq\":[],\"grantable_roles_at_other\":[],\"roles\":[],\"roles_at_base\":[],\"roles_at_hq\":[],\"roles_at_other\":[]}]")
		case "https://esi.evetech.net/v5/corporations/98627127?datasource=tranquility":
			bodyBytes = []byte("{\"ceo_id\":95067057,\"creator_id\":2115636466,\"date_founded\":\"2020-01-09T17:27:50Z\",\"description\":\"Enter a description of your corporation here.\",\"home_station_id\":60011386,\"member_count\":11,\"name\":\"Feynman Electrodynamics\",\"shares\":1000,\"tax_rate\":0.0,\"ticker\":\"FYDYN\",\"url\":\"http:\\/\\/\"}")
		case "https://esi.evetech.net/v4/corporations/98627127/members/?datasource=tranquility":
			bodyBytes = []byte("[95281762,2115692519,2115417359,95067057,2115636466,2114367476,2113199519,2115448095,2114908444,2115714045,2115692575]")
		case "https://esi.evetech.net/v3/universe/names/":
			bodyBytes = []byte("[{\"category\":\"character\",\"id\":95281762,\"name\":\"Zuberi Mwanajuma\"},{\"category\":\"character\",\"id\":2115692519,\"name\":\"Rob Barrington\"},{\"category\":\"character\",\"id\":2115417359,\"name\":\"Koriyi Chan\"},{\"category\":\"character\",\"id\":95067057,\"name\":\"Gwen Facero\"},{\"category\":\"character\",\"id\":2115636466,\"name\":\"Ion of Chios\"},{\"category\":\"character\",\"id\":2114367476,\"name\":\"Koriyo -Skill1 Skill\"},{\"category\":\"character\",\"id\":2113199519,\"name\":\"azullunes\"},{\"category\":\"character\",\"id\":2115448095,\"name\":\"Koriyo -Skill2 Skill\"},{\"category\":\"character\",\"id\":2114908444,\"name\":\"Gudrun Yassavi\"},{\"category\":\"character\",\"id\":2115714045,\"name\":\"Luke Lovell\"},{\"category\":\"character\",\"id\":2115692575,\"name\":\"Jill Kenton\"}]")
		case "https://esi.evetech.net/v1/characters/2115636466/killmails/recent/":
			bodyBytes = []byte("[]")
		case "https://esi.evetech.net/v1/corporations/98627127/killmails/recent/":
			bodyBytes = []byte("[]")

		}

		return bodyBytes, err, resp
	}
	HttpRequestMock = response

	ctrlObj.AuthCb = func(newChar *EsiChar) {
		waitForAuth <- newChar.CharInfoData.CharacterName
	}

	var newChar EsiChar
	newChar.stateMagicNum = 1628332218
	NewChar = &newChar
	start := time.Now()
	urlStrShutdown := fmt.Sprintf("http://127.0.0.1:4716/callback?code=shutdown&state=0")
	if util.SendReq(urlStrShutdown) {
		time.Sleep(400 * time.Millisecond)
	}
	elapsed := time.Since(start)
	log.Printf("UpdateGui took %s", elapsed)
	ctrlObj.StartServer()
	time.Sleep(500 * time.Millisecond)
	urlStr := fmt.Sprintf("http://127.0.0.1:4716/callback?code=your-code-here&state=1628332218")

	if !util.SendReq(urlStr) {
		t.Fatalf("request fail")
	}
	expectedChar := "Ion of Chios"
	expectedCorp := "Feynman Electrodynamics"
	select {
	case <-time.After(5000000 * time.Second):
		t.Errorf("timeout wating for auth")
	case charName := <-waitForAuth:
		if charName != expectedChar {
			t.Fatalf("unexpected character name after auth. expected %s got %s", expectedChar, charName)
		}
	}
	if len(ctrlObj.Esi.EsiCharList) == 0 {
		t.Fatalf("Charlist empty after auth")
	}
	if len(ctrlObj.Esi.EsiCorpList) == 0 {
		t.Fatalf("Corplist empty after auth")
	}
	if ctrlObj.Esi.EsiCharList[0].CharInfoData.CharacterName != expectedChar {
		t.Fatalf("unexpected character name after auth. expected %s got %s",
			expectedChar, ctrlObj.Esi.EsiCharList[0].CharInfoData.CharacterName)
	}
	if ctrlObj.Esi.EsiCorpList[0].Name != expectedCorp {
		t.Fatalf("unexpected corp name after auth. expected %s got %s",
			expectedCorp, ctrlObj.Esi.EsiCorpList[0].Name)
	}
	char := ctrlObj.Esi.EsiCharList[0]
	char.UpdateFlags.Contracts = true
	char.UpdateFlags.IndustryJobs = true
	char.UpdateFlags.Journal = true
	char.UpdateFlags.PapLinks = true
	char.UpdateFlags.Killmails = true
	char.UpdateFlags.Structures = true
	char.UpdateFlags.Assets = true

	err := ctrlObj.Save(TstCfgJson, true)
	if err != nil {
		t.Fatalf("Error writing %s", TstCfgJson)
	}
	ctrlObj.UpdateCorpMembers(char, true)
	memberIdMap := ctrlObj.Model.GetCorpMemberIdMap(char.CharInfoExt.CooperationId)
	if _, ok := memberIdMap[95281762]; !ok {
		t.Fatalf("Error character not found")
	}
	if !ctrlObj.ServerCancelled() {
		ctrlObj.HTTPShutdown()
	}
	modelObj.CloseDB()
}

func initTestObj(t *testing.T) *Ctrl {
	modelObj := model.NewModel(model.DbNameCtrlTest, false)
	ctrlObj := NewCtrl(modelObj)
	err := ctrlObj.Load(TstCfgJson, true)
	if err != nil {
		t.Fatalf("Error reading %s", TstCfgJson)
	}
	if len(ctrlObj.Esi.EsiCharList) == 0 {
		t.Fatalf("Charlist empty")
	}
	if len(ctrlObj.Esi.EsiCorpList) == 0 {
		t.Fatalf("corplist empty")
	}
	return ctrlObj
}

func TestContracts(t *testing.T) {
	ctrlObj := initTestObj(t)
	testCharContractTitle := "testTitle char " + fmt.Sprintf("%d", time.Now().Unix())
	testCorpContractTitle := "testTitle corp " + fmt.Sprintf("%d", time.Now().Unix())
	expiresOn := util.UnixTS2DateTimeStr(time.Now().Add(1199 * time.Second).Unix())
	testCharContractID := 172994018
	testCorpContractPrice := float64(300000000000.0)
	testCharContractPrice := float64(10000000.0)
	testContractOutstanding := fmt.Sprintf(`
			{
				"acceptor_id": 0,
				"assignee_id": 0,
				"availability": "public",
				"collateral": 0.0,
				"contract_id": %d,
				"date_expired": "2021-08-20T17:49:17Z",
				"date_issued": "2021-07-23T17:49:17Z",
				"days_to_complete": 0,
				"end_location_id": 60003760,
				"for_corporation": false,
				"issuer_corporation_id": 98627127,
				"issuer_id": 2115636466,
				"price": %3.2f,
				"reward": 0.0,
				"start_location_id": 60000000,
				"status": "outstanding",
				"title": "%s",
				"type": "item_exchange",
				"volume": 0.01
			}`, testCharContractID, testCharContractPrice, testCharContractTitle)
	testContractFinished := fmt.Sprintf(`
			{
				"acceptor_id": 95465499,
				"assignee_id": 0,
				"availability": "public",
				"collateral": 0.0,
				"contract_id": %d,
				"date_accepted":"2021-07-17T16:22:59Z",
				"date_completed":"2021-07-17T16:22:59Z",
				"date_expired": "2021-08-20T17:49:17Z",
				"date_issued": "2021-07-23T17:49:17Z",
				"days_to_complete": 0,
				"end_location_id": 60000000,
				"for_corporation": false,
				"issuer_corporation_id": 98627127,
				"issuer_id": 2115636466,
				"price": %3.2f,
				"reward": 0.0,
				"start_location_id": 60000000,
				"status": "finished",
				"title": "%s",
				"type": "item_exchange",
				"volume": 0.01
			}`, testCharContractID, testCharContractPrice, testCharContractTitle)
	testCorpContractID := 173494284
	testContractCorp := fmt.Sprintf(`
			{
				"acceptor_id": 95465499,
				"assignee_id": 0,
				"availability": "public",
				"collateral": 0.0,
				"contract_id": %d,
				"date_accepted": "2021-08-08T21:40:38Z",
				"date_completed": "2021-08-08T21:40:38Z",
				"date_expired": "2021-09-05T17:20:46Z",
				"date_issued": "2021-08-08T17:20:46Z",
				"days_to_complete": 0,
				"end_location_id": 1000000000001,
				"for_corporation": true,
				"issuer_corporation_id": 98627127,
				"issuer_id": 2115636466,
				"price": %3.2f,
				"reward": 0.0,
				"start_location_id": 1000000000001,
				"status": "finished",
				"title": "%s",
				"type": "item_exchange",
				"volume": 0.0
			}`, testCorpContractID, testCorpContractPrice, testCorpContractTitle)

	HttpRequestMock = func(req *http.Request) (bodyBytes []byte, err error, resp *http.Response) {
		resp = &http.Response{
			StatusCode: http.StatusNotFound,
		}
		switch req.URL.String() {
		case "https://login.eveonline.com/v2/oauth/token":
			bodyBytes = []byte(`{
									"access_token": "access_token-dummy-token",
									"expires_in": 1199,
									"token_type": "Bearer",
									"refresh_token": "refresh_token_dummytoken"
								}`)
			resp.StatusCode = http.StatusOK
		case "https://login.eveonline.com/oauth/verify":
			resultString := fmt.Sprintf(`{
													"CharacterID": 2115636466,
													"CharacterName": "Ion of Chios",
													"ExpiresOn": "%s",
													"Scopes": "publicData esi-wallet.read_character_wallet.v1 esi-wallet.read_corporation_wallet.v1 esi-universe.read_structures.v1 esi-killmails.read_killmails.v1 esi-corporations.read_corporation_membership.v1 esi-corporations.read_structures.v1 esi-industry.read_character_jobs.v1 esi-contracts.read_character_contracts.v1 esi-killmails.read_corporation_killmails.v1 esi-corporations.track_members.v1 esi-wallet.read_corporation_wallets.v1 esi-characters.read_notifications.v1 esi-contracts.read_corporation_contracts.v1 esi-corporations.read_starbases.v1 esi-industry.read_corporation_jobs.v1",
													"TokenType": "Character",
													"CharacterOwnerHash": "dummyhash=",
													"IntellectualProperty": "EVE"
												}`,
				expiresOn)
			bodyBytes = []byte(resultString)
			resp.StatusCode = http.StatusOK
		case "https://esi.evetech.net/v1/characters/2115636466/contracts/?datasource=tranquility&page=1":
			bodyBytes = []byte(fmt.Sprintf("[%s]", testContractOutstanding))
			resp.StatusCode = http.StatusOK
		case "https://esi.evetech.net/v1/corporations/98627127/contracts/?datasource=tranquility&page=1":
			bodyBytes = []byte(fmt.Sprintf("[%s]", testContractCorp))
			resp.StatusCode = http.StatusOK
		case fmt.Sprintf("https://esi.evetech.net/v1/characters/2115636466/contracts/%d/items/", testCharContractID):
			bodyBytes = []byte(`[	{
										"is_included": true,
										"is_singleton": false,
										"quantity": 1,
										"record_id": 3594873402,
										"type_id": 21026
									}	]`)
			resp.StatusCode = http.StatusOK
		case "https://esi.evetech.net/v1/corporations/98627127/contracts/173494284/items/":
			bodyBytes = []byte(`[    {
										"is_included": true,
										"is_singleton": false,
										"quantity": 1,
										"record_id": 3594873401,
										"type_id": 40340
									}    ]`)
			resp.StatusCode = http.StatusOK
		}

		return bodyBytes, err, resp
	}

	char := ctrlObj.Esi.EsiCharList[0]
	if !char.UpdateFlags.Contracts {
		t.Fatalf("contract flag should be written by previous test")
	}
	ctrlObj.UpdateContracts(char, false)
	ctrList := ctrlObj.Model.GetContractsByIssuerId(char.CharInfoData.CharacterID, false)
	if len(ctrList) == 0 {
		t.Fatalf("contract list is empty!")
	}
	if util.HumanizeNumber(ctrList[0].Price) != "10m" {
		t.Errorf("unexpected price expected %s got %s", util.HumanizeNumber(ctrList[0].Price), "10m")
	}
	ctrTitle, result := ctrlObj.Model.GetStringEntry(ctrList[0].Title)
	if !result {
		t.Fatalf("could not find string entry")
	}
	if ctrTitle != testCharContractTitle {
		t.Fatalf("unexpected contract title got %s expected %s", ctrTitle, testCharContractTitle)
	}
	if ctrList[0].Status != model.Cntr_Stat_outstanding {
		t.Fatalf("unexpected contract title status %s expected %s",
			ctrlObj.Model.ContractStatusInt2Str(ctrList[0].Status),
			ctrlObj.Model.ContractStatusInt2Str(model.Cntr_Stat_outstanding))
	}
	items := ctrlObj.Model.GetContrItems(testCharContractID)
	if len(items) == 0 {
		t.Fatalf("contract item list is empty!")
	} else {
		expItem := "Capital Jump Drive Blueprint"
		itemType := ctrlObj.Model.GetTypeString(items[0].Type_id)
		if itemType != expItem {
			t.Errorf("contract item unexpected name: expected %s got %s", expItem, itemType)
		}
	}
	// test corporation contract
	ctrlObj.UpdateContracts(char, true)
	ctrList = ctrlObj.Model.GetContractsByIssuerId(char.CharInfoExt.CooperationId, true)
	if len(ctrList) == 0 {
		t.Fatalf("contract list is empty!")
	}
	if util.HumanizeNumber(ctrList[0].Price) != "300b" {
		t.Errorf("unexpected price expected %s got %s", util.HumanizeNumber(ctrList[0].Price), "300b")
	}
	ctrTitle, result = ctrlObj.Model.GetStringEntry(ctrList[0].Title)
	if !result {
		t.Fatalf("could not find string entry")
	}
	if ctrTitle != testCorpContractTitle {
		t.Errorf("unexpected contract title got %s expected %s", ctrTitle, testCorpContractTitle)
	}
	if ctrList[0].Status != model.Cntr_Stat_finished {
		t.Errorf("unexpected contract title status %s expected %s",
			ctrlObj.Model.ContractStatusInt2Str(ctrList[0].Status),
			ctrlObj.Model.ContractStatusInt2Str(model.Cntr_Stat_finished))
	}
	items = ctrlObj.Model.GetContrItems(testCorpContractID)
	if len(items) == 0 {
		t.Fatalf("contract item list is empty!")
	} else {
		expItem := "Upwell Palatine Keepstar"
		itemType := ctrlObj.Model.GetTypeString(items[0].Type_id)
		if itemType != expItem {
			t.Errorf("contract item unexpected name: expected %s got %s", expItem, itemType)
		}
	}

	// test transition from outstanding to finished
	HttpRequestMock = func(req *http.Request) (bodyBytes []byte, err error, resp *http.Response) {
		resp = &http.Response{
			StatusCode: http.StatusNotFound,
		}
		switch req.URL.String() {
		case "https://esi.evetech.net/v1/characters/2115636466/contracts/?datasource=tranquility&page=1":
			bodyBytes = []byte(fmt.Sprintf("[%s]", testContractFinished))
			resp.StatusCode = http.StatusOK
		}
		return bodyBytes, err, resp
	}
	ctrlObj.UpdateContracts(char, false)
	ctrList = ctrlObj.Model.GetContractsByIssuerId(char.CharInfoData.CharacterID, false)
	if len(ctrList) == 0 {
		t.Fatalf("contract list is empty!")
	}
	if ctrList[0].Status != model.Cntr_Stat_finished {
		t.Fatalf("unexpected contract title status %s expected %s",
			ctrlObj.Model.ContractStatusInt2Str(ctrList[0].Status),
			ctrlObj.Model.ContractStatusInt2Str(model.Cntr_Stat_finished))
	}

	ctrlObj.Model.CloseDB()
}

func TestIndustry(t *testing.T) {
	ctrlObj := initTestObj(t)
	char := ctrlObj.Esi.EsiCharList[0]
	if !char.UpdateFlags.IndustryJobs {
		t.Fatalf("IndustryJobs flag should be written by previous test")
	}
	min := 450000000
	max := 459999999
	jobIdChar := rand.Intn((max - min) + min)
	jobIdCorp := rand.Intn((max - min) + min)
	facilityID := 1022861711365
	facilityName := "Fildar - Newark Station"
	startDateCharJob := util.UnixTS2DateTimeStr(time.Now().Unix())
	endDateCharJob := util.UnixTS2DateTimeStr(time.Now().Add(72 * time.Hour).Unix())
	startDateCorpJob := util.UnixTS2DateTimeStr(time.Now().Unix())
	endDateCorpJob := util.UnixTS2DateTimeStr(time.Now().Add(70 * time.Hour).Unix())

	testCopyJob := fmt.Sprintf(`
    {
        "activity_id": 5,
        "blueprint_id": 1000000000001,
        "blueprint_location_id": 1000000000002,
        "blueprint_type_id": 43910,
        "cost": 22818708.0,
        "duration": 2601000,
        "end_date": "%sZ",
        "facility_id": %d,
        "installer_id": 2115636466,
        "job_id": %d,
        "licensed_runs": 10,
        "output_location_id": 1000000000004,
        "probability": 1.0,
        "product_type_id": 43910,
        "runs": 50,
        "start_date": "%sZ",
        "station_id": 1000000000005,
        "status": "active"
    }`, endDateCharJob, facilityID, jobIdChar, startDateCharJob)

	testCopyJobCorp := fmt.Sprintf(`
			  {
				"activity_id": 5,
				"blueprint_id": 1027285140101,
				"blueprint_location_id": 1031112911964,
				"blueprint_type_id": 23912,
				"cost": 2874395,
				"duration": 2774400,
				"end_date": "%sZ",
				"facility_id": %d,
				"installer_id": 2115636466,
				"job_id": %d,
				"licensed_runs": 1,
				"location_id": 1024664748454,
				"output_location_id": 1031112911964,
				"probability": 1,
				"product_type_id": 23912,
				"runs": 4,
				"start_date": "%sZ",
				"status": "active"
			  }`, endDateCorpJob, facilityID, jobIdCorp, startDateCorpJob)

	structureInfo := fmt.Sprintf(`
		{
		  "name": "%s",
		  "owner_id": 98483391,
		  "position": {
			"x": -1307707374866,
			"y": -5398112443,
			"z": -1569905999189
		  },
		  "solar_system_id": 30043410,
		  "type_id": 35825
		}`, facilityName)

	HttpRequestMock = func(req *http.Request) (bodyBytes []byte, err error, resp *http.Response) {
		resp = &http.Response{
			StatusCode: http.StatusNotFound,
		}
		switch req.URL.String() {
		case "https://esi.evetech.net/v1/characters/2115636466/industry/jobs/?datasource=tranquility":
			bodyBytes = []byte(fmt.Sprintf(`[%s]`, testCopyJob))
			resp.StatusCode = http.StatusOK
		case "https://esi.evetech.net/v1/corporations/98627127/industry/jobs/?datasource=tranquility&page=1":
			bodyBytes = []byte(fmt.Sprintf(`[%s]`, testCopyJobCorp))
			resp.StatusCode = http.StatusOK
		case fmt.Sprintf("https://esi.evetech.net/v2/universe/structures/%d/?datasource=tranquility", facilityID):
			bodyBytes = []byte(fmt.Sprintf(`%s`, structureInfo))
			resp.StatusCode = http.StatusOK
		}
		return bodyBytes, err, resp
	}

	ctrlObj.UpdateIndustry(char, false)
	industryList := ctrlObj.Model.GetIndustryJobs(char.CharInfoData.CharacterID, false)
	if len(industryList) == 0 {
		t.Fatalf("industryList is empty!")
	}
	jobFaclityName := ctrlObj.Model.GetStructureNameStr(industryList[0].FacilityId)
	if jobFaclityName != facilityName {
		t.Errorf("unexpected facility name. expected %s got %s", facilityName, jobFaclityName)
	}
	if industryList[0].JobId != jobIdChar {
		t.Errorf("unexpected job id. expected %d got %d", jobIdChar, industryList[0].JobId)
	}
	ts := ctrlObj.Model.GetNextJobEndTimeStamp(char.CharInfoData.CharacterID)
	if ts != industryList[0].EndDate {
		t.Errorf("expected char job end date to be the last end date")
	}
	// test corporation
	ctrlObj.UpdateIndustry(char, true)
	industryListCorp := ctrlObj.Model.GetIndustryJobs(char.CharInfoExt.CooperationId, true)
	if len(industryListCorp) == 0 {
		t.Fatalf("industryListCorp is empty!")
	}
	if industryListCorp[0].JobId != jobIdCorp {
		t.Errorf("unexpected job id. expected %d got %d", jobIdCorp, industryListCorp[0].JobId)
	}
	jobStatus := ctrlObj.Model.JobStatusId2Str(industryListCorp[0].Status)
	expJobStatus := "active"
	if jobStatus != expJobStatus {
		t.Errorf("unexpected job status. expected %s got %s", expJobStatus, jobStatus)
	}
	jobAct := ctrlObj.Model.JobActivityId2Str(industryListCorp[0].ActivityId)
	expAct := "Copying"
	if jobAct != expAct {
		t.Errorf("unexpected job activity. expected %s got %s", expAct, jobAct)
	}
	tsCorp := ctrlObj.Model.GetNextJobEndTimeStamp(char.CharInfoData.CharacterID)
	if tsCorp != industryListCorp[0].EndDate {
		t.Errorf("expected corp job end date to be the last end date")
	}
	// test job removed from esi
	testCopyJobCorp = ""
	ctrlObj.UpdateIndustry(char, true)
	industryListCorp = ctrlObj.Model.GetIndustryJobs(char.CharInfoExt.CooperationId, true)
	if len(industryListCorp) != 0 {
		t.Fatalf("industryList is not empty!")
	}
	if !ctrlObj.Model.JobItemExist(jobIdCorp) {
		t.Errorf("expcected job Id to be still present")
	}

	ctrlObj.Model.CloseDB()
}

func TestJournal(t *testing.T) {
	ctrlObj := initTestObj(t)
	char := ctrlObj.Esi.EsiCharList[0]
	if !char.UpdateFlags.Journal {
		t.Fatalf("IndustryJobs flag should be written by previous test")
	}
	jourDate := util.UnixTS2DateTimeStr(time.Now().AddDate(0, 0, -15).Unix())
	jourDate2 := util.UnixTS2DateTimeStr(time.Now().AddDate(0, -2, -15).Unix())
	jourRefId := 19563119475
	transID := 5670996355
	transaction := fmt.Sprintf(`
		  {
			"client_id": 95465499,
			"date": "%sZ",
			"is_buy": true,
			"is_personal": true,
			"journal_ref_id": %d,
			"location_id": 60008494,
			"quantity": 7430,
			"transaction_id": %d,
			"type_id": 4247,
			"unit_price": 27160
		  }`, jourDate, jourRefId, transID)

	jourItemChar := fmt.Sprintf(`
		 {
			"amount": -201798800,
			"balance": 1000000000.022,
			"context_id": 5670996355,
			"context_id_type": "market_transaction_id",
			"date": "%sZ",
			"description": "Market escrow release",
			"first_party_id": 2115636466,
			"id": %d,
			"reason": "",
			"ref_type": "market_escrow",
			"second_party_id": 2115636466
		  }`, jourDate, jourRefId)

	amount1 := 243877.5
	boundyPrizes := fmt.Sprintf(`
			{
				"amount": %3.2f,
				"balance": 12594883384.825,
				"context_id": 30004764,
				"context_id_type": "system_id",
				"date": "%sZ",
				"description": "Ion of Chios got bounty prizes for killing pirates in 3-DMQT",
				"first_party_id": 1000125,
				"id": 19570264653,
				"reason": "24139: 2,24140: 2,24041: 1,24042: 1,24043: 1,24044: 1,24109: 1,24110: 1",
				"ref_type": "bounty_prizes",
				"second_party_id": %d,
				"tax": 243877.49999999997,
				"tax_receiver_id": %d
			  }`, amount1, jourDate, char.CharInfoData.CharacterID, char.CharInfoExt.CooperationId)
	amount2 := 127566.33
	boundyPrizes2 := fmt.Sprintf(`
			  {
				"amount": %3.2f,
				"balance": 12596348458.795,
				"context_id": 30004762,
				"context_id_type": "system_id",
				"date": "%sZ",
				"description": "Ion of Chios got bounty prizes for killing pirates in N-8YET",
				"first_party_id": 1000125,
				"id": 19571318293,
				"reason": "11906: 1,11907: 2,11908: 1,23266: 1,23250: 3,23257: 2,23258: 2,11042: 2,23267: 2,10280: 1,13691: 1,13692: 2,13693: 1,11902: 1,11903: 1",
				"ref_type": "bounty_prizes",
				"second_party_id": %d,
				"tax": 127566.32542092,
				"tax_receiver_id": %d
			  }`, amount2, jourDate2, char.CharInfoData.CharacterID, char.CharInfoExt.CooperationId)

	ECCPrizes := fmt.Sprintf(`
			  {
				"amount": 88360.2,
				"balance": 12594971745.025,
				"context_id": 30004762,
				"date": "%sZ",
				"description": "Encounter Surveillance System in N-8YET transferred funds to Ion of Chios",
				"first_party_id": 1000132,
				"id": 19570333349,
				"reason": "",
				"ref_type": "ess_escrow_transfer",
				"second_party_id": %d
			  }`, jourDate, char.CharInfoData.CharacterID)
	amount3 := 88360.2
	HttpRequestMock = func(req *http.Request) (bodyBytes []byte, err error, resp *http.Response) {
		resp = &http.Response{
			StatusCode: http.StatusNotFound,
		}
		log.Printf(req.URL.String())
		switch req.URL.String() {
		case "https://esi.evetech.net/v1/characters/2115636466/wallet/transactions/?datasource=tranquility":
			bodyBytes = []byte(fmt.Sprintf(`[%s]`, transaction))
			resp.StatusCode = http.StatusOK
		case "https://esi.evetech.net/v6/characters/2115636466/wallet/journal/?datasource=tranquility&page=1":
			bodyBytes = []byte(fmt.Sprintf(`[%s]`, jourItemChar))
			resp.StatusCode = http.StatusOK
		case fmt.Sprintf("https://esi.evetech.net/v4/corporations/%d/wallets/1/journal?datasource=tranquility&page=1", char.CharInfoExt.CooperationId):
			bodyBytes = []byte(fmt.Sprintf(`[%s, %s, %s]`, boundyPrizes, ECCPrizes, boundyPrizes2))
			resp.StatusCode = http.StatusOK
		}
		return bodyBytes, err, resp
	}
	ctrlObj.UpdateTransaction(char, false)
	ctrlObj.UpdateJournal(char, false, 0)
	journalList := ctrlObj.Model.GetJournal(char.CharInfoData.CharacterID, char.CharInfoExt.CooperationId, false)
	if len(journalList) == 0 {
		t.Fatalf("journallist is empty!")
	}
	transItem := ctrlObj.Model.GetTransactionEntry(int64(jourRefId))
	if transItem.TransactionID != int64(transID) {
		t.Errorf("unexpected transaction ID: expected %d got %d", transID, transItem.TransactionID)
	}
	// test corporation
	ctrlObj.UpdateJournal(char, true, 1)
	bounties := ctrlObj.Model.GetBounties(char.CharInfoExt.CooperationId)
	if len(bounties) == 0 {
		t.Fatalf("bountylist is empty!")
	}
	var sum float64
	for _, bounty := range bounties {
		sum += bounty.Amount
	}
	if sum != amount1+amount2+amount3 {
		t.Errorf("unexpected sum of bounties: expected %3.2f got %3.2f", sum, amount1+amount2)
	}
	bountytable := ctrlObj.Model.GetBountyTable(char.CharInfoExt.CooperationId)
	year, month, _ := time.Now().AddDate(0, 0, -15).Date()
	ymStr := fmt.Sprintf("%02d-%02d", year-2000, month)
	if val, ok := bountytable.ValCharPerMon[char.CharInfoData.CharacterName][ymStr]; ok {
		if val != amount1+amount3 {
			t.Errorf("unexpected bounties: expected %3.2f got %3.2f", amount1, val)
		}
	} else {
		t.Errorf("unexpected result no bounty found at [%s][%s]", char.CharInfoData.CharacterName, ymStr)
	}
	// test skipped values
	ctrlObj.UpdateJournal(char, true, 1)
	bounties2 := ctrlObj.Model.GetBounties(char.CharInfoExt.CooperationId)
	if len(bounties2) == 0 {
		t.Fatalf("bountylist2 is empty!")
	}
	var sum2 float64
	for _, bounty := range bounties {
		sum2 += bounty.Amount
	}
	if sum2 != amount1+amount2+amount3 {
		t.Errorf("unexpected sum of bounties: expected %3.2f got %3.2f", sum, amount1+amount2)
	}
	ctrlObj.Model.CloseDB()
}

func TestAdash(t *testing.T) {
	t.Skipf("adash to be removed")
	TestAdashFlag = true
	ctrlObj := initTestObj(t)
	char := ctrlObj.Esi.EsiCharList[0]
	if !char.UpdateFlags.PapLinks {
		t.Fatalf("PapLinks flag should be written by previous test")
	}
	time1 := util.UnixTS2AdashDateTimeStr(time.Now().AddDate(0, -1, 0).Unix())
	char1 := "Zuberi Mwanajuma"
	time2 := util.UnixTS2AdashDateTimeStr(time.Now().AddDate(0, -2, 0).Unix())
	char2 := "Rob Barrington"
	time3 := util.UnixTS2AdashDateTimeStr(time.Now().AddDate(0, -3, 0).Unix())
	char3 := "Luke Lovell"
	time4 := util.UnixTS2AdashDateTimeStr(time.Now().AddDate(0, -4, 0).Unix())
	char4 := "Jill Kenton"
	charList := []string{char1, char2, char3, char4}
	emailOrig := "test@user.com"
	pwOrig := "-&$=01#zWVr7!_dummy_password." + util.GenerateRandomString(15)

	ADhttpGetMock = func(url string, data url.Values) (bodyBytes []byte, err error, resp *http.Response) {
		resp = &http.Response{
			StatusCode: http.StatusNotFound,
		}
		switch url {
		case "https://adashboard.info/":
			bodyBytes = []byte(fmt.Sprintf(`%s`, `
					<!DOCTYPE html>
					<html lang="en">
					   <form action="/login" method="POST" class="form-horizontal">
						  <fieldset>
							 <legend><i class="glyphicons-icon lock"></i> Please Login</legend>
							 <div class="control-group ">
								<label for="Email address" class="control-label">Email address</label>
								<div class="controls">
								   <input type="text" id="Email address" name="Email address" value="" class="input-xlarge" placeholder="E-mail address"/>
								   <span class="help-inline"></span>
								</div>
							 </div>
							 <div class="control-group ">
								<label for="Password" class="control-label">Password</label>
								<div class="controls">
								   <input type="password" id="Password" name="Password" class="input-large" placeholder="Password"/>
								   <span class="help-inline"></span>
								</div>
							 </div>
							 <div class="form-actions">
								<button type="submit" value="submit" class="btn btn-primary btn-large"><i class="halflings-icon user white"></i> Log in</button>
							 </div>
						  </fieldset>
						  <small><a href="/account/forgotten">Forgotten password?</a> <a href="/account/new">Register?</a></small>
					   </form>
					   </body>
					</html>`))
			resp.StatusCode = http.StatusOK
		case "https://adashboard.info//login":
			bodyBytes = []byte(fmt.Sprintf(`%s`, `
					<!DOCTYPE html>
					<html lang="en">
						<head><meta charset="utf-8">
							<title>aD - Corporation: Feynman Electrodynamics</title>
						</head>
						<body>
						<li><a href="/corporation/FYDYN">Feynman Electrodynamics</a></li>
						</body>
					</html>`))
			userOk := false
			pwOk := false
			if data != nil {
				if value, ok := data["Email address"]; ok {
					if value[0] == emailOrig {
						userOk = true
					}
				}
				if value, ok := data["Password"]; ok {
					if value[0] == pwOrig {
						pwOk = true
					}
				}
			}
			if userOk && pwOk {
				resp.StatusCode = http.StatusOK
			}

		case "https://adashboard.info/corporation/FYDYN":
			bodyBytes = []byte(fmt.Sprintf(`%s`, `
			<!DOCTYPE html>
			<html lang="en">
			   <head>
				  <meta charset="utf-8">
				  <title>aD - Corporation: Omicron Project</title>
			   </head>
			   <body>
				  <div class="container">
					 <h5>Closed participation details</h5>
					 This is a list of participation links that have had their statistics computed.
					 This means that you are able to make aggregated statistics from these fleets. It only lists the last 25 seen.
					 <div class="accordion" id="participationAccordionWebpart1">
						<div class="accordion-group">
						   <div id="mLMNw" class="accordion-body collapse">
							  <div class="accordion-inner">
								 <p>
									<small><tt>View:</tt></small> <a href="/par/view/mLMNw">mLMNw</a><br>
								 </p>
							  </div>
						   </div>
						</div>
						<div class="accordion-group">
						   <div id="LntZv" class="accordion-body collapse">
							  <div class="accordion-inner">
								 <p>
									<small><tt>View:</tt></small> <a href="/par/view/LntZv">LntZv</a><br>
								 </p>
							  </div>
						   </div>
						</div>
						<div class="accordion-group">
						   <div id="o7C1W" class="accordion-body collapse">
							  <div class="accordion-inner">
								 <p>
									<small><tt>View:</tt></small> <a href="/par/view/o7C1W">o7C1W</a><br>
								 </p>
							  </div>
						   </div>
						</div>
						<div class="accordion-group">
						   <div id="71GhP" class="accordion-body collapse">
							  <div class="accordion-inner">
								 <p>
									<small><tt>View:</tt></small> <a href="/par/view/71GhP">71GhP</a><br>
								 </p>
							  </div>
						   </div>
						</div>
						<div class="accordion-group">
						   <div id="Y1Ifn" class="accordion-body collapse">
							  <div class="accordion-inner">
								 <p>
									<small><tt>View:</tt></small> <a href="/par/view/Y1Ifn">Y1Ifn</a><br>
								 </p>
							  </div>
						   </div>
						</div>
					 </div>
				  </div>
			   </body>
			</html>`))
			resp.StatusCode = http.StatusOK
		case "https://adashboard.info/par/export/LntZv":
			bodyBytes = []byte(fmt.Sprintf(`chName,coTicker,alShort,shTypeName,sub,loc,when
				"%s","BASTN","FYDYN","Nemesis","Facethekings","E2-RDQ","%s"`, char1, time1))
			resp.StatusCode = http.StatusOK
		case "https://adashboard.info/par/export/o7C1W":
			bodyBytes = []byte(fmt.Sprintf(`chName,coTicker,alShort,shTypeName,sub,loc,when
				"%s","BASTN","FYDYN","Nemesis","Facethekings","E2-RDQ","%s"`, char2, time2))
			resp.StatusCode = http.StatusOK
		case "https://adashboard.info/par/export/71GhP":
			bodyBytes = []byte(fmt.Sprintf(`chName,coTicker,alShort,shTypeName,sub,loc,when
				"%s","BASTN","FYDYN","Nemesis","Facethekings","E2-RDQ","%s"`, char3, time3))
			resp.StatusCode = http.StatusOK
		case "https://adashboard.info/par/export/Y1Ifn":
			bodyBytes = []byte(fmt.Sprintf(`chName,coTicker,alShort,shTypeName,sub,loc,when
				"%s","BASTN","FYDYN","Nemesis","Facethekings","E2-RDQ","%s"`, char4, time4))
			resp.StatusCode = http.StatusOK
		case "https://adashboard.info/par/export/mLMNw":
			time0 := util.UnixTS2AdashDateTimeStr(time.Now().AddDate(0, 0, 0).Unix())
			tableString := ""
			for _, charElem := range charList {
				tableString += fmt.Sprintf(`"%s","BASTN","FYDYN","Nemesis","Facethekings","E2-RDQ","%s"`+"\n", charElem, time0)
			}
			oustring := fmt.Sprintf("chName,coTicker,alShort,shTypeName,sub,loc,when\n%s", tableString)
			bodyBytes = []byte(oustring)
			resp.StatusCode = http.StatusOK

		}
		return bodyBytes, err, resp
	}
	corpID := char.CharInfoExt.CooperationId
	var email string
	var pw string

	// test overwriting existing auth by init with dummy pw
	ctrlObj.Model.SetAuth(corpID, emailOrig, "dummy")
	// now write the real auth
	ctrlObj.Model.SetAuth(corpID, emailOrig, pwOrig)

	if ctrlObj.Model.ADashAuthExists(corpID) {
		var ok bool
		email, pw, ok = ctrlObj.Model.GetAuth(corpID)
		if !ok {
			t.Fatalf("failed to decrypt password")
		}
		if email != emailOrig {
			t.Fatalf("failed to load email (expected %s got %s)", emailOrig, email)
		}
		if pw != pwOrig {
			t.Fatalf("failed to load pw (expected %s got %s)", pwOrig, pw)
		}
	} else {
		t.Fatalf("ADashAuthExists could not find corp")
	}
	ticker := ctrlObj.Model.GetCorpTicker(corpID)
	aDash := NewADashClient(email, pw, ticker, ctrlObj.Model, corpID)
	if aDash.Login() {
		if !aDash.GetPapLinks() {
			t.Fatalf("adash GetPapLinks failed")
		}
	} else {
		t.Fatalf("adash login failed")
	}
	papLinks := []string{"LntZv", "o7C1W", "71GhP", "Y1Ifn"}
	for _, link := range papLinks {
		if !ctrlObj.Model.PapLinkExists(link) {
			t.Errorf("could not find paplink %s", link)
		}
	}
	papTable := ctrlObj.Model.GetPapTable(corpID)
	if len(papTable.ValCharPerMon) == 0 {
		t.Fatalf("no entries in paptable")
	}
	year, month, _ := time.Now().AddDate(0, -3, 0).Date()
	ymStr := fmt.Sprintf("%02d-%02d", year-2000, month)
	if val, ok := papTable.ValCharPerMon[char3][ymStr]; ok {
		if val != 1 {
			t.Errorf("unexpected pap: expected %d got %d", 1, int(val))
		}
	} else {
		t.Errorf("unexpected result no entry found at [%s][%s]", char3, ymStr)
	}
	aDash.GetPapLinks()
	papTable2 := ctrlObj.Model.GetPapTable(corpID)
	var sumPaps float64
	for charname, _ := range papTable2.ValCharPerMon {
		for date, _ := range papTable2.ValCharPerMon[charname] {
			sumPaps += papTable2.ValCharPerMon[charname][date]
		}
	}
	if sumPaps != 8 {
		t.Errorf("unexpected number of paplinks expected %d got %f", 8, sumPaps)
	}
	sumPaps2 := ctrlObj.Model.GetCurrentPaps(corpID)
	if sumPaps2 != 5 {
		t.Errorf("unexpected number of paplinks expected %d got %d", 5, sumPaps2)
	}
}

func TestKillmails(t *testing.T) {
	TestAdashFlag = true
	ctrlObj := initTestObj(t)
	char := ctrlObj.Esi.EsiCharList[0]
	if !char.UpdateFlags.Killmails {
		t.Fatalf("Killmails flag should be written by previous test")
	}
	kmhash1 := "550b5296e2ef93bffa1f40f9173b9a45c11e4b71"
	kmId1 := 94455770
	kmhash2 := "cb44ce5a4df6ab0e29accf72bd9c55eab200c23b"
	kmId2 := 94452095

	HttpRequestMock = func(req *http.Request) (bodyBytes []byte, err error, resp *http.Response) {
		resp = &http.Response{
			StatusCode: http.StatusNotFound,
		}
		switch req.URL.String() {
		case fmt.Sprintf("https://esi.evetech.net/v1/corporations/%d/killmails/recent/", char.CharInfoExt.CooperationId):
			bodyBytes = []byte(fmt.Sprintf(`[
						  {
							"killmail_hash": "%s",
							"killmail_id": %d
						  },
						  {
							"killmail_hash": "%s",
							"killmail_id": %d
						  }]`, kmhash1, kmId1, kmhash2, kmId2))
			resp.StatusCode = http.StatusOK
		case fmt.Sprintf("https://esi.evetech.net/v1/markets/prices/?datasource=tranquility"):
			bodyBytes = []byte(fmt.Sprintf(model.MarketPrices))
			resp.StatusCode = http.StatusOK
		case fmt.Sprintf("https://esi.evetech.net/v1/killmails/%d/%s", kmId1, kmhash1):
			bodyBytes = []byte(fmt.Sprintf(model.KMEXampleData1,
				util.UnixTS2DateTimeStr(time.Now().Add(-72*time.Hour).Unix())))
			resp.StatusCode = http.StatusOK
		case fmt.Sprintf("https://esi.evetech.net/v1/killmails/%d/%s", kmId2, kmhash2):
			bodyBytes = []byte(fmt.Sprintf(model.KMEXampleData2,
				util.UnixTS2DateTimeStr(time.Now().Add(-48*time.Hour).Unix())))
			resp.StatusCode = http.StatusOK
		}
		return bodyBytes, err, resp
	}
	ctrlObj.UpdateMarket(char, true)
	tenguID := 0
	for key, value := range ctrlObj.Model.ItemIDs {
		if value == "Tengu" {
			tenguID = key
			break
		}
	}
	if tenguID == 0 {
		t.Fatalf("could not find tengu in ItemIDs")
	}
	item := ctrlObj.Model.GetMarketItem(tenguID)
	if item == nil {
		t.Fatalf("find id %d in market items", tenguID)
	}

	ctrlObj.UpdateKillMails(char, true)
	// test overwrite
	ctrlObj.UpdateKillMails(char, true)
	maxMonth := 12
	lossTable := ctrlObj.Model.GetKillTable(char.CharInfoExt.CooperationId, maxMonth, true)
	year, month, _ := time.Now().AddDate(0, 0, -2).Date()
	ymStr := fmt.Sprintf("%02d-%02d", year-2000, month)
	charName := "Koriyi Chan"
	expValue := 14982775.65
	if _, ok := lossTable.ValCharPerMon[charName]; ok {
		if value, ok2 := lossTable.ValCharPerMon[charName][ymStr]; ok2 {
			if fmt.Sprintf("%3.2f", expValue) != fmt.Sprintf("%3.2f", value) {
				t.Errorf("unexpected loss value expected %3.2f got %3.2f", expValue, value)
			}
		} else {
			t.Errorf("could not find date in loss table %s", ymStr)
		}
	} else {
		t.Errorf("could not find character in loss table %s", charName)
	}
	killtable := ctrlObj.Model.GetKillTable(char.CharInfoExt.CooperationId, maxMonth, false)
	killcount := 0
	for keyChar, _ := range killtable.ValCharPerMon {
		for range killtable.ValCharPerMon[keyChar] {
			//fmt.Printf("%s %s %3.2f\n", keyChar, keyDate, value)
			killcount++
		}
	}
	if killcount != 2 {
		t.Errorf("unexpected kills got %d expected 2", killcount)
	}

	kmlist := ctrlObj.Model.GetKillsMails()
	if len(kmlist) != 2 {
		t.Fatalf("expected to entries in km list")
	}
	kmtestId := 94455770
	km := ctrlObj.Model.GetKillsMail(kmtestId)
	if km == nil {
		t.Fatalf("km not found")
	}
	expectedValue := ctrlObj.Model.GetKillValue(kmtestId)
	if km.Value == 0 {
		t.Fatalf("km value not correct ")
	}
	if fmt.Sprintf("%3.2f", km.Value) != fmt.Sprintf("%3.2f", expectedValue) {
		t.Fatalf("km value not correct expected %3.2f got %3.2f", km.Value, expectedValue)
	}
	if km.Killmail_id != int32(kmtestId) {
		t.Fatalf("km id not correct")
	}

}

func TestNotifications(t *testing.T) {
	TestAdashFlag = true
	ctrlObj := initTestObj(t)
	char := ctrlObj.Esi.EsiCharList[0]
	HttpRequestMock = func(req *http.Request) (bodyBytes []byte, err error, resp *http.Response) {
		resp = &http.Response{
			StatusCode: http.StatusNotFound,
		}
		switch req.URL.String() {
		case fmt.Sprintf("https://esi.evetech.net/v6/characters/%d/notifications/?datasource=tranquility", char.CharInfoData.CharacterID):
			bodyBytes = []byte(fmt.Sprintf(`[
					  {
						"notification_id": 1339650850,
						"sender_id": 1000137,
						"sender_type": "corporation",
						"text": "solarsystemID: 30001725\nstructureID: &id001 1024664748454\nstructureShowInfoData:\n- showinfo\n- 35825\n- *id001\nstructureTypeID: 35825\n",
						"timestamp": "2020-11-25T20:07:00Z",
						"type": "StructureWentHighPower"
					  },
					  {
						"notification_id": 1339340975,
						"sender_id": 1000137,
						"sender_type": "corporation",
						"text": "solarsystemID: 30001725\nstructureID: &id001 1024664748454\nstructureShowInfoData:\n- showinfo\n- 35825\n- *id001\nstructureTypeID: 35825\n",
						"timestamp": "2020-11-25T06:02:00Z",
						"type": "StructureWentLowPower"
					  },
					  {
						"notification_id": 1339340458,
						"sender_id": 1000137,
						"sender_type": "corporation",
						"text": "isCorpOwned: true\nsolarsystemID: 30001725\nstructureID: &id001 1024664748454\nstructureLink: <a href=\"showinfo:35825//1024664748454\">Nosodnis - Nara</a>\nstructureShowInfoData:\n- showinfo\n- 35825\n- *id001\nstructureTypeID: 35825\n",
						"timestamp": "2020-11-25T06:00:00Z",
						"type": "StructuresJobsPaused"
					  },
					  {
						"notification_id": 1339340431,
						"sender_id": 1000137,
						"sender_type": "corporation",
						"text": "listOfServiceModuleIDs:\n- 35891\nsolarsystemID: 30001725\nstructureID: &id001 1024664748454\nstructureShowInfoData:\n- showinfo\n- 35825\n- *id001\nstructureTypeID: 35825\n",
						"timestamp": "2020-11-25T06:00:00Z",
						"type": "StructureServicesOffline"
					  },
					  {
						"notification_id": 1338871202,
						"sender_id": 1000137,
						"sender_type": "corporation",
						"text": "listOfTypesAndQty:\n- - 211\n  - 4247\nsolarsystemID: 30001725\nstructureID: &id001 1024664748454\nstructureShowInfoData:\n- showinfo\n- 35825\n- *id001\nstructureTypeID: 35825\n",
						"timestamp": "2020-11-24T06:48:00Z",
						"type": "StructureFuelAlert"
					  },
					  {
						"notification_id": 1191368523,
						"sender_id": 1000137,
						"sender_type": "corporation",
						"text": "listOfServiceModuleIDs:\n- 35891\nsolarsystemID: 30001725\nstructureID: &id001 1024664748454\nstructureShowInfoData:\n- showinfo\n- 35825\n- *id001\nstructureTypeID: 35825\n",
						"timestamp": "2020-03-15T05:00:00Z",
						"type": "StructureServicesOffline"
					  },
					  {
						"notification_id": 1190761965,
						"sender_id": 1000137,
						"sender_type": "corporation",
						"text": "listOfTypesAndQty:\n- - 212\n  - 4247\nsolarsystemID: 30001725\nstructureID: &id001 1024664748454\nstructureShowInfoData:\n- showinfo\n- 35825\n- *id001\nstructureTypeID: 35825\n",
						"timestamp": "2020-03-14T05:47:00Z",
						"type": "StructureFuelAlert"
					  },
					  {
						"notification_id": 1334853756,
						"sender_id": 1000137,
						"sender_type": "corporation",
						"text": "allianceID: null\ncorpID: 98548313\nmoonID: 40269956\nsolarSystemID: 30004263\ntypeID: 20060\nwants:\n- quantity: 236\n  typeID: 4247\n",
						"timestamp": "2020-11-16T07:37:00Z",
						"type": "TowerResourceAlertMsg"
					  },
					  {
						"notification_id": 1325699554,
						"sender_id": 1000023,
						"sender_type": "corporation",
						"text": "amount: 4531905\nbillTypeID: 2\ncreditorID: 1000023\ncurrentDate: 132484783635427690\ndebtorID: 98548313\ndueDate: 132510701370000000\nexternalID: 27\nexternalID2: 60002662\n",
						"timestamp": "2020-10-29T20:53:00Z",
						"type": "CorpAllBillMsg"
					  },
					{
						"is_read": true,
						"notification_id": 1322047730,
						"sender_id": 1000125,
						"sender_type": "corporation",
						"text": "againstID: 98548313\ndeclaredByID: 98659355\nendDate: 132480132000000000\n",
						"timestamp": "2020-10-23T11:40:00Z",
						"type": "WarRetractedByConcord"
					},
					  {
						"is_read": true,
						"notification_id": 1318854340,
						"sender_id": 1000137,
						"sender_type": "corporation",
						"text": "solarsystemID: 30001725\nstructureID: &id001 1024664748454\nstructureShowInfoData:\n- showinfo\n- 35825\n- *id001\nstructureTypeID: 35825\ntimeLeft: 1672097806579\ntimestamp: 132476044830000000\nvulnerableTime: 9000000000\n",
						"timestamp": "2020-10-17T19:41:00Z",
						"type": "StructureLostShields"
					  },
					  {
						"notification_id": 1318834105,
						"sender_id": 1000137,
						"sender_type": "corporation",
						"text": "allianceID: null\narmorPercentage: 100.0\ncharID: 2115802471\ncorpLinkData:\n- showinfo\n- 2\n- 98659355\ncorpName: The Inner Monastery\nhullPercentage: 100.0\nshieldPercentage: 94.93374528147335\nsolarsystemID: 30001725\nstructureID: &id001 1024664748454\nstructureShowInfoData:\n- showinfo\n- 35825\n- *id001\nstructureTypeID: 35825\n",
						"timestamp": "2020-10-17T19:08:00Z",
						"type": "StructureUnderAttack"
					  },
					  {
						"is_read": true,
						"notification_id": 1318008494,
						"sender_id": 1000125,
						"sender_type": "corporation",
						"text": "againstID: 98548313\ncost: 100000000\ndeclaredByID: 98659355\ndelayHours: 24\nhostileState: false\ntimeStarted: 132474057000000000\nwarHQ: <b>Assiad - Chicken Nuggies</b>\nwarHQ_IdType:\n- 1034571818419\n- 35832\n",
						"timestamp": "2020-10-16T10:56:00Z",
						"type": "WarDeclared"
					  },
					  {
						"notification_id": 1319408648,
						"sender_id": 95281762,
						"sender_type": "character",
						"text": "applicationText: ''\ncharID: 95281762\ncorpID: 98627127\n",
						"timestamp": "2020-10-18T17:10:00Z",
						"type": "CharAppAcceptMsg"
					  }	]`))
			resp.StatusCode = http.StatusOK
		}
		return bodyBytes, err, resp
	}

	ctrlObj.UpdateNotifications(char, false)
	notifies := ctrlObj.Model.GetCharNotifications(char.CharInfoData.CharacterID)
	if len(notifies) == 0 {
		t.Fatalf("no notifications found")
	}
	foundAttack := false
	for _, noti := range notifies {
		if noti.Type == model.NotiMsgTyp_StructureUnderAttack {
			foundAttack = true
			IDexpected := int64(1318834105)
			if noti.NotificationId != IDexpected {
				t.Errorf("unexpected Notification ID got %d expected %d", noti.NotificationId, IDexpected)
			}
			break
		}
	}
	if !foundAttack {
		t.Errorf("attack notification not found!")
	}
}

func TestStructures(t *testing.T) {
	TestAdashFlag = true
	ctrlObj := initTestObj(t)
	char := ctrlObj.Esi.EsiCharList[0]
	if !char.UpdateFlags.Structures {
		t.Fatalf("Structures flag should be written by previous test")
	}
	strucName := "Barkrik - Red Dwarf"
	solarSystemID := 30002071
	structureInfo := fmt.Sprintf(`
		{
		  "name": "%s",
		  "owner_id": 98627127,
		  "position": {
			"x": -1307707374866,
			"y": -5398112443,
			"z": -1569905999189
		  },
		  "solar_system_id": %d,
		  "type_id": 35825
		}`, strucName, solarSystemID)

	structureId := int64(1024665740000) + int64(rand.Intn((10000-1000)+1000))
	expiresOn := util.UnixTS2DateTimeStr(time.Now().AddDate(0, 0, 5).Unix())
	HttpRequestMock = func(req *http.Request) (bodyBytes []byte, err error, resp *http.Response) {
		resp = &http.Response{
			StatusCode: http.StatusNotFound,
		}
		switch req.URL.String() {

		case fmt.Sprintf("https://esi.evetech.net/v2/universe/structures/%d/?datasource=tranquility", structureId):
			bodyBytes = []byte(fmt.Sprintf(`%s`, structureInfo))
			resp.StatusCode = http.StatusOK

		case fmt.Sprintf("https://esi.evetech.net/v4/corporations/%d/structures/?datasource=tranquility", char.CharInfoExt.CooperationId):

			bodyBytes = []byte(fmt.Sprintf(`[
					  {
						"corporation_id": 98627127,
						"fuel_expires": "%sZ",
						"name": "%s",
						"profile_id": 84013,
						"reinforce_hour": 20,
						"services": [
						  {
							"name": "Material Efficiency Research",
							"state": "online"
						  },
						  {
							"name": "Blueprint Copying",
							"state": "online"
						  },
						  {
							"name": "Reprocessing",
							"state": "offline"
						  },
						  {
							"name": "Manufacturing (Standard)",
							"state": "offline"
						  },
						  {
							"name": "Time Efficiency Research",
							"state": "online"
						  }
						],
						"state": "shield_vulnerable",
						"structure_id": %d,
						"system_id": %d,
						"type_id": 35825
					  }
					]`, expiresOn, strucName, structureId, solarSystemID))
			resp.StatusCode = http.StatusOK
		}
		return bodyBytes, err, resp
	}
	ctrlObj.UpdateStructures(char, true)
	structureList := ctrlObj.Model.GetCorpStructures(char.CharInfoExt.CooperationId)
	if len(structureList) == 0 {
		t.Fatalf("could not find structures")
	}

	svcMapping := make(map[int64][]*model.DBstructureService)
	testOnlineOK := false
	testOfflineOK := false
	for _, structure := range structureList {
		name := ctrlObj.Model.GetStructureNameStr(structure.StructureID)
		if name != strucName {
			t.Errorf("unexpected struc name. expected %s got %s ", strucName, name)
		}
		strSvcs := ctrlObj.Model.GetServiceEntries(structure.StructureID)
		svcMapping[structure.StructureID] = strSvcs
		for _, svc := range strSvcs {

			svcName, _ := ctrlObj.Model.GetStringEntry(svc.Name)
			svcState, _ := ctrlObj.Model.GetStringEntry(svc.State)
			if svcName == "Material Efficiency Research" {
				if svcState == "online" {
					testOnlineOK = true
				}
			}
			if svcName == "Reprocessing" {
				if svcState == "offline" {
					testOfflineOK = true
				}
			}
		}
	}
	if !testOnlineOK {
		t.Errorf("online test failed")
	}
	if !testOfflineOK {
		t.Errorf("online test failed")
	}

	HttpRequestMock = func(req *http.Request) (bodyBytes []byte, err error, resp *http.Response) {
		resp = &http.Response{
			StatusCode: http.StatusNotFound,
		}
		switch req.URL.String() {

		case fmt.Sprintf("https://esi.evetech.net/v4/corporations/%d/structures/?datasource=tranquility", char.CharInfoExt.CooperationId):

			bodyBytes = []byte(fmt.Sprintf(`[
					  {
						"corporation_id": 98627127,
						"fuel_expires": "%sZ",
						"name": "%s",
						"profile_id": 84013,
						"reinforce_hour": 19,
						"services": [
						  {
							"name": "Material Efficiency Research",
							"state": "offline"
						  },
						  {
							"name": "Blueprint Copying",
							"state": "online"
						  },
						  {
							"name": "Reprocessing",
							"state": "online"
						  },
						  {
							"name": "Time Efficiency Research",
							"state": "online"
						  }
						],
						"state": "shield_vulnerable",
						"structure_id": %d,
						"system_id": %d,
						"type_id": 35825
					  }
					]`, expiresOn, strucName, structureId, solarSystemID))
			resp.StatusCode = http.StatusOK
		}
		return bodyBytes, err, resp
	}

	ctrlObj.UpdateStructures(char, true)
	structureList = ctrlObj.Model.GetCorpStructures(char.CharInfoExt.CooperationId)
	if len(structureList) == 0 {
		t.Fatalf("could not find structures")
	}

	svcMapping = make(map[int64][]*model.DBstructureService)
	testOnlineOK = false
	testOfflineOK = false
	for _, structure := range structureList {
		if structure.ReinForceHour != 19 {
			t.Errorf("unexpcted reinforce hour expcted %d got %d", 19, structure.ReinForceHour)
		}
		name := ctrlObj.Model.GetStructureNameStr(structure.StructureID)
		if name != strucName {
			t.Errorf("unexpected struc name. expected %s got %s ", strucName, name)
		}
		strSvcs := ctrlObj.Model.GetServiceEntries(structure.StructureID)
		svcMapping[structure.StructureID] = strSvcs
		for _, svc := range strSvcs {

			svcName, _ := ctrlObj.Model.GetStringEntry(svc.Name)
			svcState, _ := ctrlObj.Model.GetStringEntry(svc.State)
			if svcName == "Material Efficiency Research" {
				if svcState == "offline" {
					testOnlineOK = true
				}
			}
			if svcName == "Reprocessing" {
				if svcState == "online" {
					testOfflineOK = true
				}
			}
		}
	}
	if !testOnlineOK {
		t.Errorf("online test failed")
	}
	if !testOfflineOK {
		t.Errorf("online test failed")
	}
	status := ctrlObj.Model.GetStructureStatus(structureId)
	if status != "shield_vulnerable" {
		t.Errorf("unexpected structure status %s", status)
	}

}

func TestWallet(t *testing.T) {
	ctrlObj := initTestObj(t)
	char := ctrlObj.Esi.EsiCharList[0]
	exampleWalletChar := 7395847.57
	exampleWalletCorp := 100000000.0
	HttpRequestMock = func(req *http.Request) (bodyBytes []byte, err error, resp *http.Response) {
		resp = &http.Response{
			StatusCode: http.StatusNotFound,
		}
		switch req.URL.String() {
		case fmt.Sprintf("https://esi.evetech.net/v1/corporations/%d/wallets/?datasource=tranquility", char.CharInfoExt.CooperationId):
			bodyBytes = []byte(fmt.Sprintf(`[{"balance":%3.2f,"division":1},{"balance":0.0,"division":2},{"balance":0.0,"division":3},{"balance":0.0,"division":4},{"balance":0.0,"division":5},{"balance":0.0,"division":6},{"balance":0.0,"division":7}]`, exampleWalletCorp))
			resp.StatusCode = http.StatusOK
		case fmt.Sprintf("https://esi.evetech.net/v1/characters/%d/wallet/?datasource=tranquility", char.CharInfoData.CharacterID):
			bodyBytes = []byte(fmt.Sprintf(`%3.2f`, exampleWalletChar))
			resp.StatusCode = http.StatusOK
		}
		return bodyBytes, err, resp
	}
	if ctrlObj.Model.WalletEntryExists(0, char.CharInfoExt.CooperationId, 1) {
		t.Errorf("unexpected corp wallet")
	}
	if ctrlObj.Model.WalletEntryExists(char.CharInfoData.CharacterID, 0, 0) {
		t.Errorf("unexpected char wallet")
	}
	ctrlObj.UpdateWallet(char, false)
	ctrlObj.UpdateWallet(char, true)
	walletChar := ctrlObj.Model.GetLatestWallets(char.CharInfoData.CharacterID, 0, 0)
	if fmt.Sprintf(`%3.2f`, exampleWalletChar) != fmt.Sprintf(`%3.2f`, walletChar) {
		t.Errorf("unexpected wallet balance %s got %s", fmt.Sprintf(`%3.2f`, exampleWalletChar), fmt.Sprintf(`%3.2f`, walletChar))
	}
	walletcorp := ctrlObj.Model.GetLatestWallets(0, char.CharInfoExt.CooperationId, 1)
	if fmt.Sprintf(`%3.2f`, exampleWalletCorp) != fmt.Sprintf(`%3.2f`, walletcorp) {
		t.Errorf("unexpected wallet balance %s got %s", fmt.Sprintf(`%3.2f`, exampleWalletCorp), fmt.Sprintf(`%3.2f`, walletChar))
	}
	exampleWalletChar = 100.12
	exampleWalletCorp = 110000000.0
	time.Sleep(1 * time.Second)
	ctrlObj.UpdateWallet(char, false)
	ctrlObj.UpdateWallet(char, true)
	walletChar = ctrlObj.Model.GetLatestWallets(char.CharInfoData.CharacterID, 0, 0)
	exampleWalletCharStr := fmt.Sprintf(`%3.2f`, exampleWalletChar)
	walletCharStr := fmt.Sprintf(`%3.2f`, walletChar)
	if exampleWalletCharStr != walletCharStr {
		t.Errorf("unexpected wallet balance %s got %s", exampleWalletCharStr, walletCharStr)
	}
	walletcorp = ctrlObj.Model.GetLatestWallets(0, char.CharInfoExt.CooperationId, 1)
	exampleWalletCorpStr := fmt.Sprintf(`%3.2f`, exampleWalletCorp)
	walletcorpStr := fmt.Sprintf(`%3.2f`, walletcorp)
	if exampleWalletCorpStr != walletcorpStr {
		t.Errorf("unexpected wallet balance %s got %s", exampleWalletCorpStr, walletcorpStr)
	}
	if !ctrlObj.Model.WalletEntryExists(0, char.CharInfoExt.CooperationId, 1) {
		t.Errorf("could not find corp wallet")
	}
	if !ctrlObj.Model.WalletEntryExists(char.CharInfoData.CharacterID, 0, 0) {
		t.Errorf("could not find char wallet")
	}
}

func TestAssets(t *testing.T) {
	ctrlObj := initTestObj(t)
	expiresOn := util.UnixTS2DateTimeStr(time.Now().Add(1199 * time.Second).Unix())
	testAssetSingletonId := 1042491487363
	testAssetSingletonLocation := 1031112911964
	testAssetSingletonLFlag := "CorpSAG3"
	testAssetSingletonTID := 19725


	testAssetId := 1028132445071
	testAssetLocation := 1031700940306
	testAssetLFlag := "Unlocked"
	testAssetTID := 24558
	testAssetQ:=2


	HttpRequestMock = func(req *http.Request) (bodyBytes []byte, err error, resp *http.Response) {
		resp = &http.Response{
			StatusCode: http.StatusNotFound,
		}
		switch req.URL.String() {
		case "https://login.eveonline.com/v2/oauth/token":
			bodyBytes = []byte(`{
									"access_token": "access_token-dummy-token",
									"expires_in": 1199,
									"token_type": "Bearer",
									"refresh_token": "refresh_token_dummytoken"
								}`)
			resp.StatusCode = http.StatusOK
		case "https://login.eveonline.com/oauth/verify":
			resultString := fmt.Sprintf(`{
													"CharacterID": 2115636466,
													"CharacterName": "Ion of Chios",
													"ExpiresOn": "%s",
													"Scopes": "publicData esi-wallet.read_character_wallet.v1 esi-wallet.read_corporation_wallet.v1 esi-universe.read_structures.v1 esi-killmails.read_killmails.v1 esi-corporations.read_corporation_membership.v1 esi-corporations.read_structures.v1 esi-industry.read_character_jobs.v1 esi-contracts.read_character_contracts.v1 esi-killmails.read_corporation_killmails.v1 esi-corporations.track_members.v1 esi-wallet.read_corporation_wallets.v1 esi-characters.read_notifications.v1 esi-contracts.read_corporation_contracts.v1 esi-corporations.read_starbases.v1 esi-industry.read_corporation_jobs.v1",
													"TokenType": "Character",
													"CharacterOwnerHash": "dummyhash=",
													"IntellectualProperty": "EVE"
												}`,
				expiresOn)
			bodyBytes = []byte(resultString)
			resp.StatusCode = http.StatusOK
		case "https://esi.evetech.net/v5/characters/2115636466/assets/?datasource=tranquility&page=1":
			testAssetSingleton := fmt.Sprintf(`
  {
    "is_blueprint_copy": true,
    "is_singleton": true,
    "item_id": %d,
    "location_flag": "%s",
    "location_id": %d,
    "location_type": "item",
    "quantity": 1,
    "type_id": %d
  }
`, testAssetSingletonId, testAssetSingletonLFlag, testAssetSingletonLocation, testAssetSingletonTID)
			bodyBytes = []byte(fmt.Sprintf("[%s]", testAssetSingleton))
			resp.StatusCode = http.StatusOK
		case "https://esi.evetech.net/v5/corporations/98627127/assets/?datasource=tranquility&page=1":
			testAsset := fmt.Sprintf(`
  {
    "is_singleton": false,
    "item_id": %d,
    "location_flag": "%s",
    "location_id": %d,
    "location_type": "item",
    "quantity": %d,
    "type_id": %d
  }
`, testAssetId, testAssetLFlag, testAssetLocation, testAssetQ, testAssetTID)
			bodyBytes = []byte(fmt.Sprintf("[%s]", testAsset))
			resp.StatusCode = http.StatusOK
		}

		return bodyBytes, err, resp
	}

	char := ctrlObj.Esi.EsiCharList[0]
	if !char.UpdateFlags.Assets {
		t.Fatalf("Assets flag should be written by previous test")
	}
	ctrlObj.UpdateAssets(char, false)
}
