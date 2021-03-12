// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package wellarchitectediface provides an interface to enable mocking the AWS Well-Architected Tool service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package wellarchitectediface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/wellarchitected"
)

// WellArchitectedAPI provides an interface to enable mocking the
// wellarchitected.WellArchitected service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Well-Architected Tool.
//    func myFunc(svc wellarchitectediface.WellArchitectedAPI) bool {
//        // Make svc.AssociateLenses request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := wellarchitected.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockWellArchitectedClient struct {
//        wellarchitectediface.WellArchitectedAPI
//    }
//    func (m *mockWellArchitectedClient) AssociateLenses(input *wellarchitected.AssociateLensesInput) (*wellarchitected.AssociateLensesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockWellArchitectedClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type WellArchitectedAPI interface {
	AssociateLenses(*wellarchitected.AssociateLensesInput) (*wellarchitected.AssociateLensesOutput, error)
	AssociateLensesWithContext(aws.Context, *wellarchitected.AssociateLensesInput, ...request.Option) (*wellarchitected.AssociateLensesOutput, error)
	AssociateLensesRequest(*wellarchitected.AssociateLensesInput) (*request.Request, *wellarchitected.AssociateLensesOutput)

	CreateMilestone(*wellarchitected.CreateMilestoneInput) (*wellarchitected.CreateMilestoneOutput, error)
	CreateMilestoneWithContext(aws.Context, *wellarchitected.CreateMilestoneInput, ...request.Option) (*wellarchitected.CreateMilestoneOutput, error)
	CreateMilestoneRequest(*wellarchitected.CreateMilestoneInput) (*request.Request, *wellarchitected.CreateMilestoneOutput)

	CreateWorkload(*wellarchitected.CreateWorkloadInput) (*wellarchitected.CreateWorkloadOutput, error)
	CreateWorkloadWithContext(aws.Context, *wellarchitected.CreateWorkloadInput, ...request.Option) (*wellarchitected.CreateWorkloadOutput, error)
	CreateWorkloadRequest(*wellarchitected.CreateWorkloadInput) (*request.Request, *wellarchitected.CreateWorkloadOutput)

	CreateWorkloadShare(*wellarchitected.CreateWorkloadShareInput) (*wellarchitected.CreateWorkloadShareOutput, error)
	CreateWorkloadShareWithContext(aws.Context, *wellarchitected.CreateWorkloadShareInput, ...request.Option) (*wellarchitected.CreateWorkloadShareOutput, error)
	CreateWorkloadShareRequest(*wellarchitected.CreateWorkloadShareInput) (*request.Request, *wellarchitected.CreateWorkloadShareOutput)

	DeleteWorkload(*wellarchitected.DeleteWorkloadInput) (*wellarchitected.DeleteWorkloadOutput, error)
	DeleteWorkloadWithContext(aws.Context, *wellarchitected.DeleteWorkloadInput, ...request.Option) (*wellarchitected.DeleteWorkloadOutput, error)
	DeleteWorkloadRequest(*wellarchitected.DeleteWorkloadInput) (*request.Request, *wellarchitected.DeleteWorkloadOutput)

	DeleteWorkloadShare(*wellarchitected.DeleteWorkloadShareInput) (*wellarchitected.DeleteWorkloadShareOutput, error)
	DeleteWorkloadShareWithContext(aws.Context, *wellarchitected.DeleteWorkloadShareInput, ...request.Option) (*wellarchitected.DeleteWorkloadShareOutput, error)
	DeleteWorkloadShareRequest(*wellarchitected.DeleteWorkloadShareInput) (*request.Request, *wellarchitected.DeleteWorkloadShareOutput)

	DisassociateLenses(*wellarchitected.DisassociateLensesInput) (*wellarchitected.DisassociateLensesOutput, error)
	DisassociateLensesWithContext(aws.Context, *wellarchitected.DisassociateLensesInput, ...request.Option) (*wellarchitected.DisassociateLensesOutput, error)
	DisassociateLensesRequest(*wellarchitected.DisassociateLensesInput) (*request.Request, *wellarchitected.DisassociateLensesOutput)

	GetAnswer(*wellarchitected.GetAnswerInput) (*wellarchitected.GetAnswerOutput, error)
	GetAnswerWithContext(aws.Context, *wellarchitected.GetAnswerInput, ...request.Option) (*wellarchitected.GetAnswerOutput, error)
	GetAnswerRequest(*wellarchitected.GetAnswerInput) (*request.Request, *wellarchitected.GetAnswerOutput)

	GetLensReview(*wellarchitected.GetLensReviewInput) (*wellarchitected.GetLensReviewOutput, error)
	GetLensReviewWithContext(aws.Context, *wellarchitected.GetLensReviewInput, ...request.Option) (*wellarchitected.GetLensReviewOutput, error)
	GetLensReviewRequest(*wellarchitected.GetLensReviewInput) (*request.Request, *wellarchitected.GetLensReviewOutput)

	GetLensReviewReport(*wellarchitected.GetLensReviewReportInput) (*wellarchitected.GetLensReviewReportOutput, error)
	GetLensReviewReportWithContext(aws.Context, *wellarchitected.GetLensReviewReportInput, ...request.Option) (*wellarchitected.GetLensReviewReportOutput, error)
	GetLensReviewReportRequest(*wellarchitected.GetLensReviewReportInput) (*request.Request, *wellarchitected.GetLensReviewReportOutput)

	GetLensVersionDifference(*wellarchitected.GetLensVersionDifferenceInput) (*wellarchitected.GetLensVersionDifferenceOutput, error)
	GetLensVersionDifferenceWithContext(aws.Context, *wellarchitected.GetLensVersionDifferenceInput, ...request.Option) (*wellarchitected.GetLensVersionDifferenceOutput, error)
	GetLensVersionDifferenceRequest(*wellarchitected.GetLensVersionDifferenceInput) (*request.Request, *wellarchitected.GetLensVersionDifferenceOutput)

	GetMilestone(*wellarchitected.GetMilestoneInput) (*wellarchitected.GetMilestoneOutput, error)
	GetMilestoneWithContext(aws.Context, *wellarchitected.GetMilestoneInput, ...request.Option) (*wellarchitected.GetMilestoneOutput, error)
	GetMilestoneRequest(*wellarchitected.GetMilestoneInput) (*request.Request, *wellarchitected.GetMilestoneOutput)

	GetWorkload(*wellarchitected.GetWorkloadInput) (*wellarchitected.GetWorkloadOutput, error)
	GetWorkloadWithContext(aws.Context, *wellarchitected.GetWorkloadInput, ...request.Option) (*wellarchitected.GetWorkloadOutput, error)
	GetWorkloadRequest(*wellarchitected.GetWorkloadInput) (*request.Request, *wellarchitected.GetWorkloadOutput)

	ListAnswers(*wellarchitected.ListAnswersInput) (*wellarchitected.ListAnswersOutput, error)
	ListAnswersWithContext(aws.Context, *wellarchitected.ListAnswersInput, ...request.Option) (*wellarchitected.ListAnswersOutput, error)
	ListAnswersRequest(*wellarchitected.ListAnswersInput) (*request.Request, *wellarchitected.ListAnswersOutput)

	ListAnswersPages(*wellarchitected.ListAnswersInput, func(*wellarchitected.ListAnswersOutput, bool) bool) error
	ListAnswersPagesWithContext(aws.Context, *wellarchitected.ListAnswersInput, func(*wellarchitected.ListAnswersOutput, bool) bool, ...request.Option) error

	ListLensReviewImprovements(*wellarchitected.ListLensReviewImprovementsInput) (*wellarchitected.ListLensReviewImprovementsOutput, error)
	ListLensReviewImprovementsWithContext(aws.Context, *wellarchitected.ListLensReviewImprovementsInput, ...request.Option) (*wellarchitected.ListLensReviewImprovementsOutput, error)
	ListLensReviewImprovementsRequest(*wellarchitected.ListLensReviewImprovementsInput) (*request.Request, *wellarchitected.ListLensReviewImprovementsOutput)

	ListLensReviewImprovementsPages(*wellarchitected.ListLensReviewImprovementsInput, func(*wellarchitected.ListLensReviewImprovementsOutput, bool) bool) error
	ListLensReviewImprovementsPagesWithContext(aws.Context, *wellarchitected.ListLensReviewImprovementsInput, func(*wellarchitected.ListLensReviewImprovementsOutput, bool) bool, ...request.Option) error

	ListLensReviews(*wellarchitected.ListLensReviewsInput) (*wellarchitected.ListLensReviewsOutput, error)
	ListLensReviewsWithContext(aws.Context, *wellarchitected.ListLensReviewsInput, ...request.Option) (*wellarchitected.ListLensReviewsOutput, error)
	ListLensReviewsRequest(*wellarchitected.ListLensReviewsInput) (*request.Request, *wellarchitected.ListLensReviewsOutput)

	ListLensReviewsPages(*wellarchitected.ListLensReviewsInput, func(*wellarchitected.ListLensReviewsOutput, bool) bool) error
	ListLensReviewsPagesWithContext(aws.Context, *wellarchitected.ListLensReviewsInput, func(*wellarchitected.ListLensReviewsOutput, bool) bool, ...request.Option) error

	ListLenses(*wellarchitected.ListLensesInput) (*wellarchitected.ListLensesOutput, error)
	ListLensesWithContext(aws.Context, *wellarchitected.ListLensesInput, ...request.Option) (*wellarchitected.ListLensesOutput, error)
	ListLensesRequest(*wellarchitected.ListLensesInput) (*request.Request, *wellarchitected.ListLensesOutput)

	ListLensesPages(*wellarchitected.ListLensesInput, func(*wellarchitected.ListLensesOutput, bool) bool) error
	ListLensesPagesWithContext(aws.Context, *wellarchitected.ListLensesInput, func(*wellarchitected.ListLensesOutput, bool) bool, ...request.Option) error

	ListMilestones(*wellarchitected.ListMilestonesInput) (*wellarchitected.ListMilestonesOutput, error)
	ListMilestonesWithContext(aws.Context, *wellarchitected.ListMilestonesInput, ...request.Option) (*wellarchitected.ListMilestonesOutput, error)
	ListMilestonesRequest(*wellarchitected.ListMilestonesInput) (*request.Request, *wellarchitected.ListMilestonesOutput)

	ListMilestonesPages(*wellarchitected.ListMilestonesInput, func(*wellarchitected.ListMilestonesOutput, bool) bool) error
	ListMilestonesPagesWithContext(aws.Context, *wellarchitected.ListMilestonesInput, func(*wellarchitected.ListMilestonesOutput, bool) bool, ...request.Option) error

	ListNotifications(*wellarchitected.ListNotificationsInput) (*wellarchitected.ListNotificationsOutput, error)
	ListNotificationsWithContext(aws.Context, *wellarchitected.ListNotificationsInput, ...request.Option) (*wellarchitected.ListNotificationsOutput, error)
	ListNotificationsRequest(*wellarchitected.ListNotificationsInput) (*request.Request, *wellarchitected.ListNotificationsOutput)

	ListNotificationsPages(*wellarchitected.ListNotificationsInput, func(*wellarchitected.ListNotificationsOutput, bool) bool) error
	ListNotificationsPagesWithContext(aws.Context, *wellarchitected.ListNotificationsInput, func(*wellarchitected.ListNotificationsOutput, bool) bool, ...request.Option) error

	ListShareInvitations(*wellarchitected.ListShareInvitationsInput) (*wellarchitected.ListShareInvitationsOutput, error)
	ListShareInvitationsWithContext(aws.Context, *wellarchitected.ListShareInvitationsInput, ...request.Option) (*wellarchitected.ListShareInvitationsOutput, error)
	ListShareInvitationsRequest(*wellarchitected.ListShareInvitationsInput) (*request.Request, *wellarchitected.ListShareInvitationsOutput)

	ListShareInvitationsPages(*wellarchitected.ListShareInvitationsInput, func(*wellarchitected.ListShareInvitationsOutput, bool) bool) error
	ListShareInvitationsPagesWithContext(aws.Context, *wellarchitected.ListShareInvitationsInput, func(*wellarchitected.ListShareInvitationsOutput, bool) bool, ...request.Option) error

	ListWorkloadShares(*wellarchitected.ListWorkloadSharesInput) (*wellarchitected.ListWorkloadSharesOutput, error)
	ListWorkloadSharesWithContext(aws.Context, *wellarchitected.ListWorkloadSharesInput, ...request.Option) (*wellarchitected.ListWorkloadSharesOutput, error)
	ListWorkloadSharesRequest(*wellarchitected.ListWorkloadSharesInput) (*request.Request, *wellarchitected.ListWorkloadSharesOutput)

	ListWorkloadSharesPages(*wellarchitected.ListWorkloadSharesInput, func(*wellarchitected.ListWorkloadSharesOutput, bool) bool) error
	ListWorkloadSharesPagesWithContext(aws.Context, *wellarchitected.ListWorkloadSharesInput, func(*wellarchitected.ListWorkloadSharesOutput, bool) bool, ...request.Option) error

	ListWorkloads(*wellarchitected.ListWorkloadsInput) (*wellarchitected.ListWorkloadsOutput, error)
	ListWorkloadsWithContext(aws.Context, *wellarchitected.ListWorkloadsInput, ...request.Option) (*wellarchitected.ListWorkloadsOutput, error)
	ListWorkloadsRequest(*wellarchitected.ListWorkloadsInput) (*request.Request, *wellarchitected.ListWorkloadsOutput)

	ListWorkloadsPages(*wellarchitected.ListWorkloadsInput, func(*wellarchitected.ListWorkloadsOutput, bool) bool) error
	ListWorkloadsPagesWithContext(aws.Context, *wellarchitected.ListWorkloadsInput, func(*wellarchitected.ListWorkloadsOutput, bool) bool, ...request.Option) error

	UpdateAnswer(*wellarchitected.UpdateAnswerInput) (*wellarchitected.UpdateAnswerOutput, error)
	UpdateAnswerWithContext(aws.Context, *wellarchitected.UpdateAnswerInput, ...request.Option) (*wellarchitected.UpdateAnswerOutput, error)
	UpdateAnswerRequest(*wellarchitected.UpdateAnswerInput) (*request.Request, *wellarchitected.UpdateAnswerOutput)

	UpdateLensReview(*wellarchitected.UpdateLensReviewInput) (*wellarchitected.UpdateLensReviewOutput, error)
	UpdateLensReviewWithContext(aws.Context, *wellarchitected.UpdateLensReviewInput, ...request.Option) (*wellarchitected.UpdateLensReviewOutput, error)
	UpdateLensReviewRequest(*wellarchitected.UpdateLensReviewInput) (*request.Request, *wellarchitected.UpdateLensReviewOutput)

	UpdateShareInvitation(*wellarchitected.UpdateShareInvitationInput) (*wellarchitected.UpdateShareInvitationOutput, error)
	UpdateShareInvitationWithContext(aws.Context, *wellarchitected.UpdateShareInvitationInput, ...request.Option) (*wellarchitected.UpdateShareInvitationOutput, error)
	UpdateShareInvitationRequest(*wellarchitected.UpdateShareInvitationInput) (*request.Request, *wellarchitected.UpdateShareInvitationOutput)

	UpdateWorkload(*wellarchitected.UpdateWorkloadInput) (*wellarchitected.UpdateWorkloadOutput, error)
	UpdateWorkloadWithContext(aws.Context, *wellarchitected.UpdateWorkloadInput, ...request.Option) (*wellarchitected.UpdateWorkloadOutput, error)
	UpdateWorkloadRequest(*wellarchitected.UpdateWorkloadInput) (*request.Request, *wellarchitected.UpdateWorkloadOutput)

	UpdateWorkloadShare(*wellarchitected.UpdateWorkloadShareInput) (*wellarchitected.UpdateWorkloadShareOutput, error)
	UpdateWorkloadShareWithContext(aws.Context, *wellarchitected.UpdateWorkloadShareInput, ...request.Option) (*wellarchitected.UpdateWorkloadShareOutput, error)
	UpdateWorkloadShareRequest(*wellarchitected.UpdateWorkloadShareInput) (*request.Request, *wellarchitected.UpdateWorkloadShareOutput)

	UpgradeLensReview(*wellarchitected.UpgradeLensReviewInput) (*wellarchitected.UpgradeLensReviewOutput, error)
	UpgradeLensReviewWithContext(aws.Context, *wellarchitected.UpgradeLensReviewInput, ...request.Option) (*wellarchitected.UpgradeLensReviewOutput, error)
	UpgradeLensReviewRequest(*wellarchitected.UpgradeLensReviewInput) (*request.Request, *wellarchitected.UpgradeLensReviewOutput)
}

var _ WellArchitectedAPI = (*wellarchitected.WellArchitected)(nil)