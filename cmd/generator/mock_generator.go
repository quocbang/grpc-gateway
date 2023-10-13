package generator

// use mockery to generate mock file
// install mock tool
//   - go install github.com/vektra/mockery/v2@latest

// REPOSITORIES
//go:generate mockery --name Repositories --with-expecter --filename mock_repo.go --dir ../../server/repositories --output ../../server/repositories/mock
//go:generate mockery --name Transactions --with-expecter --filename mock_transaction.go --dir ../../server/repositories --output ../../server/repositories/mock
//go:generate mockery --name Account --with-expecter --filename mock_account.go --dir ../../server/repositories --output ../../server/repositories/mock

// SENDER
//go:generate mockery --name Sender --with-expecter --filename mock_sender.go --dir ../../server/sender --output ../../server/sender/mock
//go:generate mockery --name EmailInterface --with-expecter --filename mock_email_interface.go --dir ../../server/sender --output ../../server/sender/mock
