package main

import (
	"bloodBank/model"
	"bloodBank/service"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var ser = service.Connection{}

func init() {
	ser.Server = "mongodb://localhost:27017"
	ser.Database = "BloodBank"
	ser.Collection1 = "User"
	ser.Collection2 = "Donor"
	ser.Collection3 = "AvailableBlood"
	ser.Collection4 = "Patient"

	ser.Connect()
}

func saveUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var dataBody model.User
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.SaveUserDetails(dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusBadRequest, "Data saved Successfully!!!", true, result)
	}
}

func searchUsersById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		respondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	if result, err := ser.SearchUsersDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusBadRequest, "Data fetched Successfully!!!", true, result)
	}
}

func updateUserById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody model.User
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.UpdateUserDetailsById(dataBody, id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusBadRequest, "Data fetched Successfully!!!", true, result)
	}
}

func deleteUserById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	if result, err := ser.DeleteUserDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusBadRequest, "Data deleted Successfully!!!", true, result)
	}
}

func saveDonor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var donorData model.Donor
	if err := json.NewDecoder(r.Body).Decode(&donorData); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.SaveDonorData(donorData); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusBadRequest, "Data saved Successfully!!!", true, result)
	}
}

func searchDonorById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		respondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	if result, err := ser.SearchDonorDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusBadRequest, "Data fetched Successfully!!!", true, result)
	}
}

func updateDonorById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody model.Donor
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.UpdateDonorDetailsById(dataBody, id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusBadRequest, "Data updated Successfully!!!", true, result)
	}
}

func deleteDonorById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	if result, err := ser.DeleteDonorDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusBadRequest, "Data deleted Successfully!!!", true, result)
	}
}

func bloodRequestPatient(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var dataBody model.Patient
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.ApplyBloodPatientDetails(dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusBadRequest, "Data fetched Successfully!!!", true, result)
	}
}

func bloodProvidedPatient(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	segment := strings.Split(r.URL.Path, "/")
	id := segment[len(segment)-1]
	if id == "" {
		respondWithError(w, http.StatusBadRequest, "Please provide Id for Search")
	}

	if result, err := ser.GivenBloodPatientDetailsById(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusBadRequest, "Data fetched Successfully!!!", true, result)
	}
}

func searchFilterBloodDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var bloodDetailsRequest model.AvailableBlood
	if err := json.NewDecoder(r.Body).Decode(&bloodDetailsRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	if result, err := ser.SearchFilterBloodDetails(bloodDetailsRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusBadRequest, "Data fetched Successfully!!!", true, result)
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, msg, false, map[string]string{})
}

func respondWithJson(w http.ResponseWriter, code int, message string, success bool, payload interface{}) {

	resp := model.Response{Data: payload, Success: success, SuccessMsg: message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp.ToJson())
}

func main() {
	http.HandleFunc("/save-user-details", saveUser)
	http.HandleFunc("/search-user-details-id/", searchUsersById)
	http.HandleFunc("/update-user-details-id/", updateUserById)
	http.HandleFunc("/delete-user-details-id/", deleteUserById)
	http.HandleFunc("/save-donor-details", saveDonor)
	http.HandleFunc("/search-donor-details-id/", searchDonorById)
	http.HandleFunc("/update-donor-details-id/", updateDonorById)
	http.HandleFunc("/delete-donor-details-id/", deleteDonorById)
	http.HandleFunc("/blood-request-patient-details", bloodRequestPatient)
	http.HandleFunc("/blood-provided-patient-details-id/", bloodProvidedPatient)
	http.HandleFunc("/search-filter-blood-details/", searchFilterBloodDetails)
	http.HandleFunc("/generate-token", GetJwt)
	http.Handle("/api", ValidateJWT(saveDonor))
	log.Println("Server started at 8080")
	http.ListenAndServe(":8080", nil)
}

func GetJwt(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	requestBody := make(map[string]string, 2)

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	userId := requestBody["username"]
	password := requestBody["password"]

	err := ser.AuthenticateUser(password, userId)

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := CreateJWT(userId)
	if err != nil {
		return
	}
	setTokenInHeader(w, token)
	respondWithJson(w, http.StatusAccepted, "User authenticated successfully", true, map[string]string{"Tokenid": token})
}

func setTokenInHeader(w http.ResponseWriter, token string) {
	w.Header().Add("Tokenid", token)
}

var SECRET = []byte("super-secret-auth-key")

type JWTClaim struct {
	UserId string `json:"usedId"`
	jwt.StandardClaims
}

func CreateJWT(usedId string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &JWTClaim{
		UserId: usedId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Tokenid"] != nil {
			token, err := jwt.ParseWithClaims(r.Header["Token"][0], &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return SECRET, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}
			claims, ok := token.Claims.(*JWTClaim)
			if !ok {
				err = errors.New("couldn't parse claims")
				return
			}
			if claims.ExpiresAt < time.Now().Local().Unix() {
				err = errors.New("token expired")
				return
			}
			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret area")
}
