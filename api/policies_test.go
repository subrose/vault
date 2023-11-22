package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
	_vault "github.com/subrose/vault"
)

func TestPolicies(t *testing.T) {
	app, core := InitTestingVault(t)

	testPolicyId := "test-policy"

	t.Run("can create policy", func(t *testing.T) {
		testPolicy := _vault.Policy{
			PolicyId:  testPolicyId,
			Effect:    _vault.EffectAllow,
			Actions:   []_vault.PolicyAction{_vault.PolicyActionRead},
			Resources: []string{fmt.Sprintf("/policies/%s", testPolicyId)},
		}

		request := newRequest(t, http.MethodPost, "/policies", map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, testPolicy)

		response := performRequest(t, app, request)
		var createdPolicy _vault.Policy
		checkResponse(t, response, http.StatusCreated, &createdPolicy)

		// Assertions
		assert.Equal(t, _vault.EffectAllow, createdPolicy.Effect)
		assert.Equal(t, []_vault.PolicyAction{_vault.PolicyActionRead}, createdPolicy.Actions)
		assert.Equal(t, []string{fmt.Sprintf("/policies/%s", testPolicyId)}, createdPolicy.Resources)
	})

	t.Run("can get policy", func(t *testing.T) {
		request := newRequest(t, http.MethodGet, fmt.Sprintf("/policies/%s", testPolicyId), map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, nil)

		response := performRequest(t, app, request)
		var returnedPolicy _vault.Policy
		checkResponse(t, response, http.StatusOK, &returnedPolicy)

		// Assertions
		assert.Equal(t, _vault.EffectAllow, returnedPolicy.Effect)
		assert.Equal(t, []_vault.PolicyAction{_vault.PolicyActionRead}, returnedPolicy.Actions)
		assert.Equal(t, []string{fmt.Sprintf("/policies/%s", testPolicyId)}, returnedPolicy.Resources)
	})

	t.Run("can delete policy", func(t *testing.T) {
		// Add a dummy policy first before deleting
		dummyPolicy := _vault.Policy{
			PolicyId:  "dummy-policy",
			Effect:    _vault.EffectAllow,
			Actions:   []_vault.PolicyAction{_vault.PolicyActionRead},
			Resources: []string{"/policies/dummy-policy"},
		}

		request := newRequest(t, http.MethodPost, "/policies", map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, dummyPolicy)

		response := performRequest(t, app, request)
		var returnedPolicy _vault.Policy
		checkResponse(t, response, http.StatusCreated, &returnedPolicy)

		// Delete it
		request = newRequest(t, http.MethodDelete, fmt.Sprintf("/policies/%s", dummyPolicy.PolicyId), map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, nil)

		response = performRequest(t, app, request)
		checkResponse(t, response, http.StatusNoContent, nil)

		// Check it's gone
		request = newRequest(t, http.MethodGet, fmt.Sprintf("/policies/%s", dummyPolicy.PolicyId), map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, nil)

		response = performRequest(t, app, request)
		checkResponse(t, response, http.StatusNotFound, nil)
	})
}