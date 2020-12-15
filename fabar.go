/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

type SellItem struct {
	Name		string 		`json:"name"`
	Seller		string		`json:"seller"`
	Price		int64		`json:"price"`
}

type CheckRecv struct {
	RecvNum		string		`json:"recvnum"`
	BuyerName	string		`json:"buyername"`

}
type ReqEval struct {
	Gradename	string		`json:"gradename"`
	Recvnum		string		`json:"recvnum`
}

type EvalRecv struct {
	GradeNo		string		`json:"gradeno"`
	RecvNo		string		`json:"recvno"`
	GradeEval	int64		`json:"gradeeval"`
}

type BuyPurchase struct{
	RecvNum		string		`json:"recvnum"`
	BuyerName	string		`json:"buyername"`
}

type EnterOrg struct {
	OrgName		string		`json:"orgname"`
	Token		int64		`json:"token"`
	Class		string		`json:"class"`
}

type SellReceipt struct {
	ItemKey		string 		`json:"ItemKey"`
	SellerName	string		`json:"sellername"`
	BuyerName	string		`json:"buyername"`
	NumProduct	int64		`json:"numproduct"`
	TotalPrice	int64		`json:"totalprice"`
	SellDate	string		`json:"selldate"`
	DueDate		string		`json:"duedate"`
}

type Receivable struct {
	ReceiptKey	string		`json:"ReceiptKey"`
	OwnerName	string		`json:"ownername"`
	HavedList	[]string	`json:"havedlist"`
	IssueRate	float64		`json:"issuerate"`
	PublishDate	string		`json:""publishdate"`
	ExpireDate	string		`json:"expiredate"`
	IsGuarantee	bool		`json:"isguarantee"`
	IsSale		bool		`json:"issale"`
}

type ReceivableRating struct {
	RecvKey		string		`json:"RecvKey"`
	GradeKey	string		`json:"GradeKey"`
	RatingPrice float64		`json:"ratingprice"`
	RatingDate	string		`json:"ratedate"`
}

type ItemResult struct {
	Key			string		`json:"Key"`
	Record		*SellItem
}

type evalRecvResult struct {
	Key			string		`json:"Key"`
	Record		*EvalRecv
}

type ItemPriceResult struct {
	Price		string
	sellerOrg	string
}

type OrgResult struct {
	Key			string		`json:"Key"`
	Record		*EnterOrg
}

type RecptResult struct {
	Key			string		`json:"Key"`
	Record		*SellReceipt
}

type RecvResult struct {
	Key			string		`json:"Key"`
	Record		*Receivable
}

type RecvRatingResult struct {
	Key			string		`json:"Key"`
	Record		*ReceivableRating
}

type CheckbuyResult struct {
	Key    		string		 `json:"Key"`
	Record 		*CheckRecv
}

type CheckReqEval struct {
	Key 		string		`json:"Key"`
	Record		*ReqEval
}

func IsKeyStruct(itemKey string,responseKey string) bool {
	return strings.Contains(responseKey, itemKey)
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	sellItems := []SellItem {
		SellItem{Name: "Chicken", Seller: "ORG1", Price: 20000},
		SellItem{Name: "Playstation5", Seller: "ORG2", Price: 500000},
		SellItem{Name: "Coffee", Seller: "ORG3" , Price: 5000},
		SellItem{Name: "Cola", Seller: "ORG4", Price: 2000},
		SellItem{Name: "Pepsi Cola", Seller: "ORG5", Price: 2000},
		SellItem{Name: "IPhone12", Seller: "ORG6", Price: 1000000},
		SellItem{Name: "Galaxy s20 fe", Seller:"ORG7", Price: 1000000},
		SellItem{Name: "Logitech G502", Seller: "ORG8", Price: 50000},
		SellItem{Name: "MacPro 16", Seller: "ORG6", Price: 3000000},
	}

	for i, item := range sellItems {
		itemAsBytes, _ := json.Marshal(item)
		err := ctx.GetStub().PutState("ITEM"+strconv.Itoa(i), itemAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	enterOrgs := []EnterOrg {
		EnterOrg{OrgName: "BBQ", Token: 100000000, Class:"Company"},
		EnterOrg{OrgName: "Sony", Token: 100000000, Class:"Company"},
		EnterOrg{OrgName: "Starbugs", Token: 100000000, Class:"Company"},
		EnterOrg{OrgName: "Cokacola", Token: 100000000, Class:"Company"},
		EnterOrg{OrgName: "Pepsi", Token: 100000000, Class:"Company"},
		EnterOrg{OrgName: "Apple", Token: 100000000, Class:"Company"},
		EnterOrg{OrgName: "Samsung", Token: 100000000, Class:"Company"},
		EnterOrg{OrgName: "Logitech", Token: 100000000, Class:"Company"},
		EnterOrg{OrgName: "KB bank", Token: 200000000, Class:"Finance"},
		EnterOrg{OrgName: "Shinhan bank", Token: 200000000, Class:"Finance"},
		EnterOrg{OrgName: "Hana bank", Token: 200000000, Class:"Finance"},
		EnterOrg{OrgName: "Woori bank", Token: 200000000, Class:"Finance"},
		EnterOrg{OrgName: "Nice", Token: 100000000, Class:"Grade"},
		EnterOrg{OrgName: "SCI", Token: 100000000, Class:"Grade"},
		EnterOrg{OrgName: "KIS", Token: 100000000, Class:"Grade"},
		EnterOrg{OrgName: "Coupang", Token: 500000000, Class:"Buyer"},
		EnterOrg{OrgName: "Naver", Token: 500000000, Class:"Buyer"},
		EnterOrg{OrgName: "Kakao", Token: 500000000, Class:"Buyer"},
		EnterOrg{OrgName: "Himart", Token: 500000000, Class:"Buyer"},
		EnterOrg{OrgName: "Danawa", Token: 500000000, Class:"Buyer"},
	}

	for i, org := range enterOrgs {
		orgAsBytes, _ := json.Marshal(org)
		err := ctx.GetStub().PutState("ORG"+strconv.Itoa(i), orgAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	selldatelist := []string{"2020-11-27", "2020-11-27", "2020-11-27", "2020-11-27", "2020-10-16", "2020-09-16", "2020-07-10", "2020-09-02", "2020-05-14"}
	duedatelist := []string{"2020-12-31", "2021-01-31", "2021-02-30", "2021-03-31", "2022-04-30", "2021-05-16", "2022-03-23", "2023-06-07", "2024-09-09"}
	receipts := []SellReceipt {
		SellReceipt{ItemKey: "ITEM1", SellerName: "ORG1", BuyerName: "ORG16", NumProduct: 10, TotalPrice: 200000, SellDate: selldatelist[0], DueDate: duedatelist[0]},
		SellReceipt{ItemKey: "ITEM2", SellerName: "ORG2", BuyerName: "ORG17", NumProduct: 10, TotalPrice: 5000000, SellDate: selldatelist[1], DueDate: duedatelist[1]},
		SellReceipt{ItemKey: "ITEM3", SellerName: "ORG3", BuyerName: "ORG18", NumProduct: 10, TotalPrice: 50000, SellDate: selldatelist[2], DueDate: duedatelist[2]},
		SellReceipt{ItemKey: "ITEM4", SellerName: "ORG4", BuyerName: "ORG19", NumProduct: 10, TotalPrice: 20000, SellDate: selldatelist[3], DueDate: duedatelist[3]},
		SellReceipt{ItemKey: "ITEM5", SellerName: "ORG5", BuyerName: "ORG20", NumProduct: 10, TotalPrice: 20000, SellDate: selldatelist[4], DueDate: duedatelist[4]},
		SellReceipt{ItemKey: "ITEM6", SellerName: "ORG6", BuyerName: "ORG16", NumProduct: 10, TotalPrice: 10000000, SellDate: selldatelist[5], DueDate: duedatelist[5]},
		SellReceipt{ItemKey: "ITEM7", SellerName: "ORG7", BuyerName: "ORG17", NumProduct: 10, TotalPrice: 10000000, SellDate: selldatelist[6], DueDate: duedatelist[6]},
		SellReceipt{ItemKey: "ITEM8", SellerName: "ORG8", BuyerName: "ORG18", NumProduct: 10, TotalPrice: 500000, SellDate: selldatelist[7], DueDate: duedatelist[7]},
		SellReceipt{ItemKey: "ITEM9", SellerName: "ORG6", BuyerName: "ORG19", NumProduct: 10, TotalPrice: 30000000, SellDate: selldatelist[8], DueDate: duedatelist[8]},
	}

	for i, receipt := range receipts {
		receiptAsBytes, _ := json.Marshal(receipt)
		err := ctx.GetStub().PutState("REPT"+strconv.Itoa(i), receiptAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	publishdatelist := []string{"2020-11-27" ,"2020-11-28", "2020-11-29", "2020-11-30", "2020-10-18", "2020-09-19", "2020-07-17", "2020-10-03", "2020-05-29"}
	expiredatelist := []string{"2020-12-31", "2021-01-31", "2021-02-30", "2021-03-31", "2022-04-30", "2021-05-16", "2022-03-23", "2023-06-07", "2024-09-09"}
	receivables := []Receivable {
		Receivable{ReceiptKey: "REPT1", OwnerName: "ORG1", HavedList: []string{}, IssueRate: 0.1, PublishDate: publishdatelist[0], ExpireDate: expiredatelist[0], IsGuarantee: false, IsSale: true},
		Receivable{ReceiptKey: "REPT2", OwnerName: "ORG9", HavedList: []string{"ORG2"}, IssueRate: 0.15, PublishDate: publishdatelist[1], ExpireDate: expiredatelist[1], IsGuarantee: true, IsSale: false},
		Receivable{ReceiptKey: "REPT3", OwnerName: "ORG3", HavedList: []string{}, IssueRate: 0.05, PublishDate: publishdatelist[2], ExpireDate: expiredatelist[2], IsGuarantee: false, IsSale: true},
		Receivable{ReceiptKey: "REPT4", OwnerName: "ORG4", HavedList: []string{}, IssueRate: 0.02, PublishDate: publishdatelist[3], ExpireDate: expiredatelist[3], IsGuarantee: false, IsSale: true},
		Receivable{ReceiptKey: "REPT5", OwnerName: "ORG5", HavedList: []string{}, IssueRate: 0.01, PublishDate: publishdatelist[4], ExpireDate: expiredatelist[4], IsGuarantee: false, IsSale: true},
		Receivable{ReceiptKey: "REPT6", OwnerName: "ORG11", HavedList: []string{"ORG6", "ORG10"}, IssueRate:0.2, PublishDate: publishdatelist[5], ExpireDate: expiredatelist[5], IsGuarantee: true, IsSale: false},
		Receivable{ReceiptKey: "REPT7", OwnerName: "ORG12", HavedList: []string{"ORG7"}, IssueRate: 0.1, PublishDate: publishdatelist[6], ExpireDate: expiredatelist[6], IsGuarantee: false, IsSale: true},
		Receivable{ReceiptKey: "REPT8", OwnerName: "ORG8", HavedList: []string{}, IssueRate: 0.02, PublishDate: publishdatelist[7], ExpireDate: expiredatelist[7], IsGuarantee: false, IsSale: true},
		Receivable{ReceiptKey: "REPT9", OwnerName: "ORG10", HavedList: []string{"ORG6"}, IssueRate: 0.3, PublishDate: publishdatelist[8], ExpireDate: expiredatelist[8], IsGuarantee: true, IsSale: true},
	}

	for i, recv := range receivables {
		recvAsBytes, _ := json.Marshal(recv)
		err := ctx.GetStub().PutState("RECV"+strconv.Itoa(i), recvAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world stat. %s", err.Error())
		}
	}

	ratingdatelist := []string{"2020-11-27" ,"2020-11-28", "2020-11-28", "2020-11-30"}
	recvRatings := []ReceivableRating {
		ReceivableRating{RecvKey: "RECV2", GradeKey: "ORG13", RatingPrice: 5125000, RatingDate: ratingdatelist[0]},
		ReceivableRating{RecvKey: "RECV6", GradeKey: "ORG14", RatingPrice: 11600000, RatingDate: ratingdatelist[1]},
		ReceivableRating{RecvKey: "RECV6", GradeKey: "ORG15", RatingPrice: 11750000, RatingDate: ratingdatelist[2]},
		ReceivableRating{RecvKey: "RECV7", GradeKey: "ORG15", RatingPrice: 11800000, RatingDate: ratingdatelist[3]},
	}

	for i, recr := range recvRatings {
		recrAsBytes, _ := json.Marshal(recr)
		err := ctx.GetStub().PutState("RECR"+strconv.Itoa(i), recrAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world stat. %s", err.Error())
		}
	}

	return nil
}

func (s *SmartContract) CreateSellItem(ctx contractapi.TransactionContextInterface, itemNumber string, name string, seller string, price int64) error {
	item := SellItem {
		Name:	name,
		Seller:	seller,
		Price: price,
	}

	itemAsBytes, _ := json.Marshal(item)

	return ctx.GetStub().PutState(itemNumber, itemAsBytes)
}

func (s *SmartContract) CreateReqPurchase(ctx contractapi.TransactionContextInterface, reqpurchaseno string, recvnum string, buyername string) error {
	req := BuyPurchase {
		RecvNum: recvnum,
		NuyerName: buyername,
	}

	reqAsBytes, _ := json.Marshal(req)

	return ctx.GetStub().PutState(reqpurchaseno, reqAsBytes)
}

func (s *SmartContract) CreateEnterOrg(ctx contractapi.TransactionContextInterface, orgNumber string, orgname string, token int64, class string) error {
	org := EnterOrg {
		OrgName: orgname,
		Token: token,
		Class: class,
	}

	orgAsBytes, _ := json.Marshal(org)

	return ctx.GetStub().PutState(orgNumber, orgAsBytes)
}

func (s *SmartContract) reqRecvEval(ctx contractapi.TransactionContextInterface, evalnum string, gradename string, recvnum string) error {
	req := ReqEval {
		Gradename: gradename,
		Recvnum: recvnum,
	}

	reqRecvEvalasBytes, _ := json.Marshal(req)

	return ctx.GetStub().PutState(evalnum, reqRecvEvalasBytes)
}

func (s *SmartContract) evaluateRec(ctx contractapi.TransactionContextInterface, gradeno string,recvnumber string, gradeEval int64) error {
	eval := EvalRecv {
		GradeNo: gradeno,
		RecvNo: recvnumber,
		GradeEval: gradeEval,
	}

	evalAsBytes, _ := json.Marshal(eval)

	return ctx.GetStub().PutState(gradeno, evalAsBytes)
}

func (s *SmartContract) CreateSellReceipt(ctx contractapi.TransactionContextInterface, reptNumber string, itemKey string, sellername string, buyername string, numproduct int64, totalprice int64, selldate string, duedate string) error {
	receipt := SellReceipt {
		ItemKey: itemKey,
		SellerName: sellername,
		BuyerName: buyername,
		NumProduct: numproduct,
		TotalPrice: totalprice,
		SellDate: selldate,
		DueDate: duedate,
	}

	receiptAsBytes, _ := json.Marshal(receipt)

	return ctx.GetStub().PutState(reptNumber, receiptAsBytes)
}

func (s *SmartContract) CreateReceivable(ctx contractapi.TransactionContextInterface, recvNumber string, receiptKey string, ownername string, havedlist []string,issuerate float64, publishdate string, expiredate string, isguarantee bool, issale bool) error {

	receivable := Receivable {
		ReceiptKey: receiptKey,
		OwnerName: ownername,
		HavedList: havedlist,
		IssueRate: issuerate,
		PublishDate: publishdate,
		ExpireDate: expiredate,
		IsGuarantee: isguarantee,
		IsSale: issale,
	}
	
	recvAsBytes, _ := json.Marshal(receivable)

	return ctx.GetStub().PutState(recvNumber, recvAsBytes)
}

func (s *SmartContract) CreateRecvRating(ctx contractapi.TransactionContextInterface, recvRatNumber string, recvKey string, gradeKey string, ratingPrice float64, ratingDate string) error {
	recr := ReceivableRating {
		RecvKey: recvKey,
		GradeKey: gradeKey,
		RatingPrice: ratingPrice,
		RatingDate: ratingDate,
	}

	recrAsBytes, _ := json.Marshal(recr)

	return ctx.GetStub().PutState(recvRatNumber, recrAsBytes)
}

func (s *SmartContract) QuerySellItem(ctx contractapi.TransactionContextInterface, itemNumber string) (*SellItem, error) {
	itemAsBytes, err := ctx.GetStub().GetState(itemNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if itemAsBytes == nil {
		return nil, fmt.Errorf("%s does not exists", itemNumber)
	}

	item := new(SellItem)
	_ = json.Unmarshal(itemAsBytes, item)

	return item, nil
}

func (s *SmartContract) QueryEnterOrg(ctx contractapi.TransactionContextInterface, orgNumber string) (*EnterOrg, error) {
	orgAsBytes, err := ctx.GetStub().GetState(orgNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if orgAsBytes == nil {
		return nil, fmt.Errorf("%s does not exists", orgNumber)
	}

	org := new(EnterOrg)
	_ = json.Unmarshal(orgAsBytes, org)

	return org, nil
}

func (s *SmartContract) QuerySellReceipt(ctx contractapi.TransactionContextInterface, reptNumber string) (*SellReceipt, error) {
	reptAsBytes, err := ctx.GetStub().GetState(reptNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if reptAsBytes == nil {
		return nil, fmt.Errorf("%s does not exists", reptNumber)
	}

	rept := new(SellReceipt)
	_ = json.Unmarshal(reptAsBytes, rept)

	return rept, nil
}

func (s *SmartContract) QueryReceivable(ctx contractapi.TransactionContextInterface, recvNumber string) (*Receivable, error) {
	recvAsBytes, err := ctx.GetStub().GetState(recvNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if recvAsBytes == nil {
		return nil, fmt.Errorf("%s does not exists", recvNumber)
	}
	recv := new(Receivable)
	_ = json.Unmarshal(recvAsBytes, recv)

	return recv, nil
}

func (s *SmartContract) QeuryRecvRating(ctx contractapi.TransactionContextInterface, recrNumber string) (*ReceivableRating, error) {
	recrAsBytes, err := ctx.GetStub().GetState(recrNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if recrAsBytes == nil {
		return nil, fmt.Errorf("%s does not exists", recrNumber)
	}
	recr := new(ReceivableRating)
	_ = json.Unmarshal(recrAsBytes, recr)

	return recr, nil
}

func (s *SmartContract) QueryAllItems(ctx contractapi.TransactionContextInterface) ([]ItemResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []ItemResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		if IsKeyStruct("ITEM",queryResponse.Key) {
			item := new(SellItem)
			_ = json.Unmarshal(queryResponse.Value, item)

			itemResult := ItemResult{Key: queryResponse.Key, Record: item}
			results = append(results, itemResult)
		}
	}
	
	return results, nil
}

func (s *SmartContract) QueryAllOrgs(ctx contractapi.TransactionContextInterface) ([]OrgResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []OrgResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		if IsKeyStruct("ORG",queryResponse.Key) {
			org := new(EnterOrg)
			_ = json.Unmarshal(queryResponse.Value, org)

			orgResult := OrgResult{Key: queryResponse.Key, Record: org}
			results = append(results, orgResult)
		}
	}

	return results, nil
}

func (s *SmartContract) QueryAllGrades(ctx contractapi.TransactionContextInterface) ([]OrgResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []OrgResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		if IsKeyStruct("ORG",queryResponse.Key) {
			org := new(EnterOrg)
			_ = json.Unmarshal(queryResponse.Value, org)
			if IsKeyStruct("Grade",org.Class) {
				orgResult := OrgResult{Key: queryResponse.Key, Record: org}
				results = append(results, orgResult)
			}
		}
	}

	return results, nil
}

func (s *SmartContract) checkRecv(ctx contractapi.TransactionContextInterface) ([]evalRecvResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []evalRecvResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		if IsKeyStruct("evalRecv",queryResponse.Key) {
			evalRecv := new(EvalRecv)
			_ = json.Unmarshal(queryResponse.Value, evalRecv)

			evalRecvResult := evalRecvResult{Key: queryResponse.Key, Record: evalRecv}
			results = append(results, evalRecvResult)
		}
	}

	return results, nil
}

func (s *SmartContract) QueryAllFinances(ctx contractapi.TransactionContextInterface) ([]OrgResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []OrgResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		if IsKeyStruct("ORG",queryResponse.Key) {
			org := new(EnterOrg)
			_ = json.Unmarshal(queryResponse.Value, org)
			if IsKeyStruct("Finance",org.Class) {
				orgResult := OrgResult{Key: queryResponse.Key, Record: org}
				results = append(results, orgResult)
			}
		}
	}

	return results, nil
}

func (s *SmartContract) QueryAllCompanys(ctx contractapi.TransactionContextInterface) ([]OrgResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []OrgResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		if IsKeyStruct("ORG",queryResponse.Key) {
			org := new(EnterOrg)
			_ = json.Unmarshal(queryResponse.Value, org)
			if IsKeyStruct("Company",org.Class) {
				orgResult := OrgResult{Key: queryResponse.Key, Record: org}
				results = append(results, orgResult)
			}
		}
	}

	return results, nil
}

func (s *SmartContract) QueryAllBuyers(ctx contractapi.TransactionContextInterface) ([]OrgResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []OrgResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		if IsKeyStruct("ORG",queryResponse.Key) {
			org := new(EnterOrg)
			_ = json.Unmarshal(queryResponse.Value, org)
			if IsKeyStruct("Buyer",org.Class) {
				orgResult := OrgResult{Key: queryResponse.Key, Record: org}
				results = append(results, orgResult)
			}
		}
	}

	return results, nil
}


func (s *SmartContract) QueryAllReceipts(ctx contractapi.TransactionContextInterface) ([]RecptResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []RecptResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}
	
		if IsKeyStruct("REPT",queryResponse.Key) {
			rept := new(SellReceipt)
			_ = json.Unmarshal(queryResponse.Value, rept)

			reptResult := RecptResult{Key: queryResponse.Key, Record: rept}
			results = append(results, reptResult)
		}
	}

	return results, nil
}

func (s *SmartContract) QueryBuyItems(ctx contractapi.TransactionContextInterface, buyerName string) ([]RecptResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []RecptResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		if IsKeyStruct("REPT", queryResponse.Key){
			rept := new(SellReceipt)
			_ = json.Unmarshal(queryResponse.Value, rept)

			if (rept.BuyerName == buyerName) {
				reptResult := RecptResult{Key: queryResponse.Key, Record: rept}
				results = append(results, reptResult)
			}
		}
	}

	func (s *SmartContract) CheckBuyRecv(ctx contractapi.TransactionContextInterface) ([]CheckbuyResult, error) {
		startKey := ""
		endKey := ""
	
		resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	
		if err != nil {
			return nil, err
		}
		defer resultsIterator.Close()
	
		results := []CheckbuyResult{}
	
		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
	
			if err != nil {
				return nil, err
			}
	
			checkRecv := new(CheckRecv)
			_ = json.Unmarshal(queryResponse.Value, checkRecv)
	
			checkbuyResult := CheckbuyResult{Key: queryResponse.Key, Record: checkRecv}
			results = append(results, checkbuyResult)
		}
	
		return results, nil
	}

	return results, nil
}

func (s *SmartContract) QuerySellItems(ctx contractapi.TransactionContextInterface, sellerName string) ([]RecptResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []RecptResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		if IsKeyStruct("REPT", queryResponse.Key){
			rept := new(SellReceipt)
			_ = json.Unmarshal(queryResponse.Value, rept)

			if (rept.SellerName == sellerName) {
				reptResult := RecptResult{Key: queryResponse.Key, Record: rept}
				results = append(results, reptResult)
			}
		}
	}

	return results, nil
}

func (s *SmartContract) CreateGrade(ctx contractapi.TransactionContextInterface, gradeNumber string, gradeName string, Token string, org string) error {
	grade := Grade{
		gradeName:  gradeName,
		Token: Token,
		org:  org,
	}

	gradeAsBytes, _ := json.Marshal(grade)

	return ctx.GetStub().PutState(gradeNumber, gradeAsBytes)
}

func (s *SmartContract) QueryAllReceivables(ctx contractapi.TransactionContextInterface) ([]RecvResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []RecvResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		if IsKeyStruct("RECV",queryResponse.Key) {
			recv := new(Receivable)
			_ = json.Unmarshal(queryResponse.Value, recv)

			recvResult := RecvResult{Key: queryResponse.Key, Record: recv}
			results = append(results, recvResult)
		}
	}

	return results, nil
}

func (s *SmartContract) QueryAllRecvRatings(ctx contractapi.TransactionContextInterface) ([]RecvRatingResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []RecvRatingResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}
		
		if IsKeyStruct("RECR",queryResponse.Key) {
			recr := new(ReceivableRating)
			_ = json.Unmarshal(queryResponse.Value, recr)

			recrResult := RecvRatingResult{Key: queryResponse.Key, Record: recr}
			results = append(results, recrResult)
		}
	}

	return results, nil
}

func (s *SmartContract) ChangeSeller(ctx contractapi.TransactionContextInterface, itemNumber string, newSeller string) error {
	item, err := s.QuerySellItem(ctx, itemNumber)

	if err != nil {
		return err
	}

	item.Seller = newSeller

	itemAsBytes, _ := json.Marshal(item)
	return ctx.GetStub().PutState(itemNumber, itemAsBytes)
}

func (s *SmartContract) ChangeRecvOwner(ctx contractapi.TransactionContextInterface, recvNumber string, newOwner string) error {
	recv, err := s.QueryReceivable(ctx, recvNumber)

	if err != nil {
		return err
	}
	recv.HavedList = append(recv.HavedList, recv.OwnerName)
	recv.OwnerName = newOwner

	recvAsBytes, _ := json.Marshal(recv)
	return ctx.GetStub().PutState(recvNumber, recvAsBytes)
}

func (s *SmartContract) UpdateTokenTransfer(ctx contractapi.TransactionContextInterface, sellerNumber string, buyerNumber string, totalprice int64) error {
	seller, err := s.QueryEnterOrg(ctx, sellerNumber)

	if err != nil {
		return err
	}

	buyer, err := s.QueryEnterOrg(ctx, buyerNumber)

	if err != nil {
		return err
	}

	if buyer.Token < totalprice {
		return fmt.Errorf("buyer token is lack totalprice")
	}
	buyer.Token = buyer.Token - totalprice
	seller.Token = seller.Token + totalprice
	buyerAsBytes, _ := json.Marshal(buyer)
	err = ctx.GetStub().PutState(buyerNumber, buyerAsBytes)
	if err != nil {
		return err
	}
	sellerAsBytes, _ := json.Marshal(seller)
	err = ctx.GetStub().PutState(sellerNumber, sellerAsBytes)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabar chaincode: %s", err.Error())
	}
}