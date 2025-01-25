package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type LoanContract struct {
	contractapi.Contract
}

type Loan struct {
	LoanID        string    `json:"loanID"`
	ApplicantName string    `json:"applicantName"`
	LoanAmount    float64   `json:"loanAmount"`
	TermMonths    int       `json:"termMonths"`
	InterestRate  float64   `json:"interestRate"`
	Outstanding   float64   `json:"outstanding"`
	Status        string    `json:"status"`
	Repayments    []float64 `json:"repayments"`
}

// TODO: Implement ApplyForLoan
func (c *LoanContract) ApplyForLoan(ctx contractapi.TransactionContextInterface, loanID, applicantName string, loanAmount float64, termMonths int, interestRate float64) error {

	existingLoan, err := ctx.GetStub().GetState(loanID)

	if err != nil {
		return fmt.Errorf("failed while checking if already a loan exists")
	}

	if existingLoan != nil {
		return fmt.Errorf("Loan already exists for the Loan ID")
	}

	newLoan := Loan{
		LoanID:        loanID,
		ApplicantName: applicantName,
		LoanAmount:    loanAmount,
		TermMonths:    termMonths,
		InterestRate:  interestRate,
		Outstanding:   loanAmount,
		Status:        "Applied",
		Repayments:    []float64{},
	}

	newLoanJson, err := json.Marshal(newLoan)

	if err != nil {
		return fmt.Errorf("failed to marshal loan details")
	}

	err = ctx.GetStub().PutState(loanID, newLoanJson)

	if err != nil {
		return fmt.Errorf("failed to create loan details")
	}

	return nil
}

// TODO: Implement ApproveLoan
func (c *LoanContract) ApproveLoan(ctx contractapi.TransactionContextInterface, loanID string, status string) error {

	if status != "approved" && status != "rejected" {
		return fmt.Errorf("Loan must be either approved or rejected")
	}

	loanBytes, err := ctx.GetStub().GetState(loanID)

	if err != nil {
		return fmt.Errorf("error while retrieving loan details")
	}

	if loanBytes == nil {
		return fmt.Errorf("loan details not found")
	}

	var loan Loan

	err = json.Unmarshal(loanBytes, &loan)

	if err != nil {
		return fmt.Errorf("error while unmarshaling the loan details")
	}

	loan.Status = status

	updatedloanBytes, err := json.Marshal(loan)

	if err != nil {
		return fmt.Errorf("failed while marshalling the loan data")
	}

	if updatedloanBytes == nil {
		return fmt.Errorf("error for updated loan details")
	}

	err = ctx.GetStub().PutState(loanID, updatedloanBytes)

	if err != nil {
		return fmt.Errorf("error while updating loan status")
	}

	return nil
}

// TODO: Implement MakeRepayment
func (c *LoanContract) MakeRepayment(ctx contractapi.TransactionContextInterface, loanID string, repaymentAmount float64) error {

	loanBytes, err := ctx.GetStub().GetState(loanID)

	if err != nil {
		return fmt.Errorf("erro while getting Loan details")
	}

	if loanBytes == nil {
		return fmt.Errorf("error while querying loan details for repayment")
	}

	var loan Loan

	err = json.Unmarshal(loanBytes, &loan)

	if err != nil {
		return fmt.Errorf("error while unmarshalling")
	}

	if loan.Status == "Closed" {
		return fmt.Errorf("loan with ID %s is already closed", loanID)
	}

	// Validate repayment amount
	if repaymentAmount <= 0 {
		return fmt.Errorf("repayment amount must be greater than zero")
	}
	if repaymentAmount > loan.Outstanding {
		return fmt.Errorf("repayment amount exceeds outstanding balance")
	}

	loan.Outstanding -= repaymentAmount
	loan.Repayments = append(loan.Repayments, repaymentAmount)

	// If the loan is fully repaid, update status
	if loan.Outstanding == 0 {
		loan.Status = "Closed"
	}

	// Serialize the updated loan data
	updatedLoanJSON, err := json.Marshal(loan)
	if err != nil {
		return fmt.Errorf("failed to marshal updated loan data: %v", err)
	}

	// Store updated loan back in the ledger
	err = ctx.GetStub().PutState(loanID, updatedLoanJSON)
	if err != nil {
		return fmt.Errorf("failed to update loan repayment: %v", err)
	}
	return nil
}

// TODO: Implement CheckLoanBalance
func (c *LoanContract) CheckLoanBalance(ctx contractapi.TransactionContextInterface, loanID string) (*Loan, error) {

	loanJSON, err := ctx.GetStub().GetState(loanID)
	if err != nil {
		return nil, fmt.Errorf("failed to read loan from world state: %v", err)
	}
	if loanJSON == nil {
		return nil, fmt.Errorf("loan with ID %s does not exist", loanID)
	}

	// Deserialize the loan data
	var loan Loan
	err = json.Unmarshal(loanJSON, &loan)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal loan data: %v", err)
	}

	// Return the loan details
	return &loan, nil

}

func main() {
	chaincode, err := contractapi.NewChaincode(new(LoanContract))
	if err != nil {
		fmt.Printf("Error creating loan chaincode: %s", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting loan chaincode: %s", err)
	}
}
