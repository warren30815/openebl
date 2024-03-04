package trade_document

import (
	"fmt"

	"github.com/openebl/openebl/pkg/bu_server/model"
	"github.com/openebl/openebl/pkg/bu_server/model/trade_document/bill_of_lading"
)

type FileBasedEBLAction string

const (
	// FileBasedEBLActionAllow is the constant for allow action
	FILE_EBL_UPDATE_DRAFT  FileBasedEBLAction = "UPDATE_DRAFT"
	FILE_EBL_AMEND         FileBasedEBLAction = "AMEND"
	FILE_EBL_REQUEST_AMEND FileBasedEBLAction = "REQUEST_AMEND"
	FILE_EBL_PRINT         FileBasedEBLAction = "PRINT"
	FILE_EBL_TRANSFER      FileBasedEBLAction = "TRANSFER"
	FILE_EBL_RETURN        FileBasedEBLAction = "RETURN"
	FILE_EBL_SURRENDER     FileBasedEBLAction = "SURRENDER"
	FILE_EBL_ACCOMPLISH    FileBasedEBLAction = "ACCOMPLISH"
)

type _AllowActionChecker func(bl *bill_of_lading.BillOfLadingPack, bu string, withDetail bool) error

var _allowActionChecker = map[FileBasedEBLAction]_AllowActionChecker{
	FILE_EBL_UPDATE_DRAFT:  IsFileEBLUpdatable,
	FILE_EBL_AMEND:         IsFileEBLAmendable,
	FILE_EBL_REQUEST_AMEND: IsFileEBLRequestAmendable,
	FILE_EBL_PRINT:         IsFileEBLPrintable,
	FILE_EBL_TRANSFER:      IsFileEBLTransferable,
	FILE_EBL_RETURN:        IsFileEBLReturnable,
	FILE_EBL_SURRENDER:     IsFileEBLSurrenderable,
	FILE_EBL_ACCOMPLISH:    IsFileEBLAccomplishable,
}

func CheckFileBasedEBLAllowAction(action FileBasedEBLAction, bl *bill_of_lading.BillOfLadingPack, bu string, withDetail bool) error {
	checker, ok := _allowActionChecker[action]
	if !ok && withDetail {
		return fmt.Errorf("%s is unknown action%w", action, model.ErrEBLActionNotAllowed)
	} else if !ok {
		return model.ErrEBLActionNotAllowed
	}

	return checker(bl, bu, withDetail)
}

func IsFileEBLUpdatable(bl *bill_of_lading.BillOfLadingPack, bu string, withDetail bool) error {
	if draft := GetDraft(bl); draft == nil || !*draft {
		if withDetail {
			return fmt.Errorf("not a draft%w", model.ErrEBLActionNotAllowed)
		}
		return model.ErrEBLActionNotAllowed
	}

	if issuer := GetIssuer(bl); issuer == nil || *issuer != bu {
		if withDetail {
			return fmt.Errorf("not the issuer%w", model.ErrEBLActionNotAllowed)
		}
		return model.ErrEBLActionNotAllowed
	}

	return nil
}

func IsFileEBLAmendable(bl *bill_of_lading.BillOfLadingPack, bu string, withDetail bool) error {
	return model.ErrEBLActionNotAllowed
}

func IsFileEBLRequestAmendable(bl *bill_of_lading.BillOfLadingPack, bu string, withDetail bool) error {
	return model.ErrEBLActionNotAllowed
}

func IsFileEBLPrintable(bl *bill_of_lading.BillOfLadingPack, bu string, withDetail bool) error {
	return model.ErrEBLActionNotAllowed
}

func IsFileEBLTransferable(bl *bill_of_lading.BillOfLadingPack, bu string, withDetail bool) error {
	return model.ErrEBLActionNotAllowed
}

func IsFileEBLReturnable(bl *bill_of_lading.BillOfLadingPack, bu string, withDetail bool) error {
	return model.ErrEBLActionNotAllowed
}

func IsFileEBLSurrenderable(bl *bill_of_lading.BillOfLadingPack, bu string, withDetail bool) error {
	return model.ErrEBLActionNotAllowed
}

func IsFileEBLAccomplishable(bl *bill_of_lading.BillOfLadingPack, bu string, withDetail bool) error {
	return model.ErrEBLActionNotAllowed
}