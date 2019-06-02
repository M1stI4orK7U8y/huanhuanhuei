package main

import (
	"strings"
	"testing"

	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/service/validate"
)

func TestGetBTCTxDetail(t *testing.T) {
	tx, err := validate.GetBtcTxDetail("aa3756a863921cd2f70b56d66c74228721c7108486a37ada48ff2d2d20be983e")
	if err != nil {
		t.Errorf("TestGetBTCTxDetail: get tx fail: %s", err.Error())
	}

	if strings.Compare(tx.Txid, "aa3756a863921cd2f70b56d66c74228721c7108486a37ada48ff2d2d20be983e") != 0 {
		t.Errorf("TestGetBTCTxDetail: txhash fail")
	}

	if len(tx.Vin) != 1 {
		t.Errorf("TestGetBTCTxDetail: vin length fail")
	}

	if len(tx.Vout) != 2 {
		t.Errorf("TestGetBTCTxDetail: vout length fail")
	}

	t.Logf("TestGetBTCTxDetail: Success")
}
