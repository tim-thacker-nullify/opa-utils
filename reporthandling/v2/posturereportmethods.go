package v2

import (
	"github.com/armosec/k8s-interface/workloadinterface"
	"github.com/armosec/opa-utils/objectsenvelopes"
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
	"github.com/armosec/opa-utils/reporthandling/results/v1/resourcesresults"
)

// Status get the overall scanning status
func (postureReport *PostureReport) Status(f *helpersv1.Filters) apis.ScanningStatus {
	return postureReport.SummaryDetails.Status(f)
}

// =========================================== List resources ====================================

// ListExcludedResources list all excluded resources IDs
func (postureReport *PostureReport) ListExcludedResources(f *helpersv1.Filters) []string {
	return postureReport.listResources(f, apis.StatusExcluded).ListExcluded()
}

// ListPassedResources list all passed resources IDs
func (postureReport *PostureReport) ListPassedResources(f *helpersv1.Filters) []string {
	return postureReport.listResources(f, apis.StatusPassed).ListPassed()
}

// ListSkippedResources list all skipped resources IDs
func (postureReport *PostureReport) ListSkippedResources(f *helpersv1.Filters) []string {
	return postureReport.listResources(f, apis.StatusSkipped).ListSkipped()
}

// ListFailedResources list all failed resources IDs
func (postureReport *PostureReport) ListFailedResources(f *helpersv1.Filters) []string {
	return postureReport.listResources(f, apis.StatusFailed).ListFailed()
}

// ListAllResources list all resources IDs. This function lists the resources IDs from the "results" and not from the "resources"
func (postureReport *PostureReport) ListAllResources(f *helpersv1.Filters) *helpersv1.AllLists {
	return postureReport.listResources(f, "")
}

func (postureReport *PostureReport) listResources(f *helpersv1.Filters, status apis.ScanningStatus) *helpersv1.AllLists {
	resources := &helpersv1.AllLists{}
	for i := range postureReport.Results {
		s := postureReport.Results[i].Status(f)
		if status == "" || s == status {
			resources.Append(s, postureReport.Results[i].GetResourceID())
		}
	}
	return resources
}

// =========================================== List Frameworks ====================================

// ListExcludedResources list all excluded resources IDs
func (postureReport *PostureReport) ListExcludedFrameworks() []string {
	return postureReport.SummaryDetails.ListFrameworks(apis.StatusExcluded).ListExcluded()
}

// ListPassedResources list all passed resources IDs
func (postureReport *PostureReport) ListPassedFrameworks() []string {
	return postureReport.SummaryDetails.ListFrameworks(apis.StatusPassed).ListPassed()
}

// ListSkippedResources list all skipped resources IDs
func (postureReport *PostureReport) ListSkippedFrameworks() []string {
	return postureReport.SummaryDetails.ListFrameworks(apis.StatusSkipped).ListSkipped()
}

// ListFailedResources list all failed resources IDs
func (postureReport *PostureReport) ListFailedFrameworks() []string {
	return postureReport.SummaryDetails.ListFrameworks(apis.StatusFailed).ListFailed()
}

// ListAllResources list all resources IDs. This function lists the resources IDs from the "results" and not from the "resources"
func (postureReport *PostureReport) ListAllFrameworks() *helpersv1.AllLists {
	return postureReport.SummaryDetails.ListFrameworks("")
}

// =========================================== List Controls ====================================

// ListExcludedResources list all excluded resources IDs
func (postureReport *PostureReport) ListExcludedControls(f *helpersv1.Filters) []string {
	return postureReport.SummaryDetails.ListControls(f, apis.StatusExcluded).ListExcluded()
}

// ListPassedResources list all passed resources IDs
func (postureReport *PostureReport) ListPassedControls(f *helpersv1.Filters) []string {
	return postureReport.SummaryDetails.ListControls(f, apis.StatusPassed).ListPassed()
}

// ListSkippedResources list all skipped resources IDs
func (postureReport *PostureReport) ListSkippedControls(f *helpersv1.Filters) []string {
	return postureReport.SummaryDetails.ListControls(f, apis.StatusSkipped).ListSkipped()
}

// ListFailedResources list all failed resources IDs
func (postureReport *PostureReport) ListFailedControls(f *helpersv1.Filters) []string {
	return postureReport.SummaryDetails.ListControls(f, apis.StatusFailed).ListFailed()
}

// ListAllResources list all resources IDs. This function lists the resources IDs from the "results" and not from the "resources"
func (postureReport *PostureReport) ListAllControls(f *helpersv1.Filters) *helpersv1.AllLists {
	return postureReport.SummaryDetails.ListControls(f, "")
}

// ==================================== Resource =============================================

// GetResource get single resource in IMetadata interface representation
func (postureReport *PostureReport) GetResource(resourceID string) workloadinterface.IMetadata {
	for i := range postureReport.Resources {
		if postureReport.Resources[i].ResourceID == resourceID {
			if m, ok := postureReport.Resources[i].Object.(map[string]interface{}); ok {
				return objectsenvelopes.NewObject(m)
			}
			break
		}
	}
	return nil
}

// ResourceStatus get single resource status. If resource not found will return an empty string
func (postureReport *PostureReport) ResourceStatus(resourceID string, f *helpersv1.Filters) apis.ScanningStatus {
	for i := range postureReport.Results {
		if postureReport.Results[i].ResourceID == resourceID {
			return postureReport.Results[i].Status(f)
		}
	}
	return ""
}

// ResourceResult get the result of a single resource. If resource not found will return nil
func (postureReport *PostureReport) ResourceResult(resourceID string) *resourcesresults.Result {
	for i := range postureReport.Results {
		if postureReport.Results[i].ResourceID == resourceID {
			return &postureReport.Results[i]
		}
	}
	return nil
}