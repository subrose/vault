package main

import (
	"fmt"
	"net/http"
	"testing"

	_vault "github.com/subrose/vault"
)

func TestCollections(t *testing.T) {
	app, core := InitTestingVault(t)

	customerCollection := CollectionModel{
		Name: "customers",
		Fields: map[string]CollectionFieldModel{
			"name":         {Type: "name", IsIndexed: true},
			"phone_number": {Type: "phoneNumber", IsIndexed: true},
			"dob":          {Type: "date", IsIndexed: false},
		},
	}

	t.Run("can create a collection", func(t *testing.T) {
		request := newRequest(t, http.MethodPost, "/collections", map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, customerCollection)

		response := performRequest(t, app, request)
		checkResponse(t, response, http.StatusCreated, nil)

	})

	t.Run("can get a collection", func(t *testing.T) {
		request := newRequest(t, http.MethodGet, "/collections/customers", map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, nil)

		response := performRequest(t, app, request)
		var returnedCollection CollectionModel
		checkResponse(t, response, http.StatusOK, &returnedCollection)

		if returnedCollection.Name != "customers" {
			t.Error("Error getting collection", returnedCollection)
		}

		if returnedCollection.Fields["name"].Type != "name" {
			t.Error("Error getting collection", returnedCollection)
		}
	})

	t.Run("can get all collections", func(t *testing.T) {
		request := newRequest(t, http.MethodGet, "/collections", map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, nil)

		response := performRequest(t, app, request)

		var returnedCollections []string
		checkResponse(t, response, http.StatusOK, &returnedCollections)
	})

	t.Run("can delete a collection", func(t *testing.T) {
		// Create a dummy collection
		collectionToDelete := CollectionModel{
			Name: "delete-me",
			Fields: map[string]CollectionFieldModel{
				"name": {Type: "name", IsIndexed: true},
			},
		}
		request := newRequest(t, http.MethodPost, "/collections", map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, collectionToDelete)

		response := performRequest(t, app, request)
		checkResponse(t, response, http.StatusCreated, nil)
		// Delete it
		request = newRequest(t, http.MethodDelete, "/collections/delete-me", map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, nil)

		response = performRequest(t, app, request)
		checkResponse(t, response, http.StatusOK, nil)
	})

	t.Run("can create and get a record", func(t *testing.T) {
		records := []map[string]interface{}{
			{
				"name":         "123345",
				"phone_number": "12345",
				"dob":          "12345",
			},
		}

		request := newRequest(t, http.MethodPost, "/collections/customers/records", map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, records)

		response := performRequest(t, app, request)
		var returnedRecordIds []string
		checkResponse(t, response, http.StatusCreated, &returnedRecordIds)
		if len(returnedRecordIds) != 1 {
			t.Error("Error creating record", returnedRecordIds)
		}
		// Get the record
		request = newRequest(t, http.MethodGet, fmt.Sprintf("/collections/customers/records/%s?formats=name.plain", returnedRecordIds[0]), map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, nil)

		response = performRequest(t, app, request)
		var returnedRecords map[string]interface{}
		checkResponse(t, response, http.StatusOK, returnedRecords)
	})

	t.Run("can update a record", func(t *testing.T) {
		// Create a record to update
		records := []map[string]interface{}{
			{
				"name":         "123345",
				"phone_number": "12345",
				"dob":          "12345",
			},
		}

		request := newRequest(t, http.MethodPost, "/collections/customers/records", map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, records)

		response := performRequest(t, app, request)
		var returnedRecordIds []string
		checkResponse(t, response, http.StatusCreated, &returnedRecordIds)
		if len(returnedRecordIds) != 1 {
			t.Error("Error creating record", returnedRecordIds)
		}

		// Update the record
		updateRecord := map[string]interface{}{
			"name":         "54321",
			"phone_number": "54321",
			"dob":          "54321",
		}

		request = newRequest(t, http.MethodPut, fmt.Sprintf("/collections/customers/records/%s", returnedRecordIds[0]), map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, updateRecord)

		response = performRequest(t, app, request)
		checkResponse(t, response, http.StatusOK, nil)

		// Get the updated record
		request = newRequest(t, http.MethodGet, fmt.Sprintf("/collections/customers/records/%s?formats=name.plain,dob.plain,phone_number.plain", returnedRecordIds[0]), map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, nil)

		response = performRequest(t, app, request)
		var returnedRecords map[string]_vault.Record
		checkResponse(t, response, http.StatusOK, &returnedRecords)

		if returnedRecords[returnedRecordIds[0]]["name"] != "54321" || returnedRecords[returnedRecordIds[0]]["phone_number"] != "54321" || returnedRecords[returnedRecordIds[0]]["dob"] != "54321" {
			t.Error("Error updating record", returnedRecords)
		}
	})

	t.Run("can delete a record", func(t *testing.T) {
		// Create a record to delete
		records := []map[string]interface{}{
			{
				"name":         "123345",
				"phone_number": "12345",
				"dob":          "12345",
			},
		}

		request := newRequest(t, http.MethodPost, "/collections/customers/records", map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, records)

		response := performRequest(t, app, request)
		var returnedRecordIds []string
		checkResponse(t, response, http.StatusCreated, &returnedRecordIds)
		if len(returnedRecordIds) != 1 {
			t.Error("Error creating record", returnedRecordIds)
		}

		// Delete the record
		request = newRequest(t, http.MethodDelete, fmt.Sprintf("/collections/customers/records/%s", returnedRecordIds[0]), map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, nil)

		response = performRequest(t, app, request)
		checkResponse(t, response, http.StatusOK, nil)

		// Try to get the deleted record
		request = newRequest(t, http.MethodGet, fmt.Sprintf("/collections/customers/records/%s?formats=name.plain", returnedRecordIds[0]), map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, nil)

		response = performRequest(t, app, request)
		checkResponse(t, response, http.StatusNotFound, nil)
	})

	t.Run("cant create a bad record", func(t *testing.T) {
		badRecords := []map[string]interface{}{
			{
				"xxx":          "123345",
				"phone_number": "12345",
				"dob":          "12345",
			},
		}

		request := newRequest(t, http.MethodPost, "/collections/customers/records", map[string]string{
			"Authorization": createBasicAuthHeader(core.conf.VAULT_ADMIN_USERNAME, core.conf.VAULT_ADMIN_PASSWORD),
		}, badRecords)

		response := performRequest(t, app, request)
		checkResponse(t, response, http.StatusBadRequest, nil)
	})

	t.Run("unauthenticated user cant crud a collection", func(t *testing.T) { // TODO: Can probably make this a table test?
		records := []map[string]interface{}{
			{
				"name":         "123345",
				"phone_number": "12345",
				"dob":          "12345",
			},
		}

		request := newRequest(t, http.MethodPost, "/collections/customers/records", map[string]string{}, records)

		response := performRequest(t, app, request)
		checkResponse(t, response, http.StatusUnauthorized, nil)

		request = newRequest(t, http.MethodGet, "/collections/customers/records/123345", map[string]string{
			"Authorization": createBasicAuthHeader("bad", "bad"),
		}, nil)
		response = performRequest(t, app, request)
		checkResponse(t, response, http.StatusUnauthorized, nil)

	})
}
