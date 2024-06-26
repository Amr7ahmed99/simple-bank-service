package tests

// func TestTransferAPI(t *testing.T) {
// 	amount := int64(10)

// 	user1 := CreateRandomUser(t)
// 	user2 := CreateRandomUser(t)
// 	user3 := CreateRandomUser(t)

// 	account1 := randomAccount(user1.Username)
// 	account2 := randomAccount(user2.Username)
// 	account3 := randomAccount(user3.Username)

// 	account1.CurrencyID = enums.USD
// 	account2.CurrencyID = enums.USD
// 	account3.CurrencyID = enums.EUR

// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        enums.USD,
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, authorizationTypeBearer, tokenMaker, user1.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(account1, nil)
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.ID)).Times(1).Return(account2, nil)

// 				arg := db.TransferTxParams{
// 					FromAccountId: account1.ID,
// 					ToAccountId:   account2.ID,
// 					Amount:        amount,
// 				}
// 				store.EXPECT().TransferTx(gomock.Any(), gomock.Eq(arg)).Times(1)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "UnauthorizedUser",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        enums.USD,
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, authorizationTypeBearer, tokenMaker, user2.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(account1, nil)
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.ID)).Times(0)
// 				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusUnauthorized, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "NoAuthorization",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        enums.USD,
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
// 				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusUnauthorized, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "FromAccountNotFound",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        enums.USD,
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, authorizationTypeBearer, tokenMaker, user1.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(db.Account{}, sql.ErrNoRows)
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.ID)).Times(0)
// 				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusNotFound, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "ToAccountNotFound",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        enums.USD,
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, authorizationTypeBearer, tokenMaker, user1.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(account1, nil)
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.ID)).Times(1).Return(db.Account{}, sql.ErrNoRows)
// 				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusNotFound, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "FromAccountCurrencyMismatch",
// 			body: gin.H{
// 				"from_account_id": account3.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        enums.USD,
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, authorizationTypeBearer, tokenMaker, user3.Username, time.Minute)

// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account3.ID)).Times(1).Return(account3, nil)
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.ID)).Times(0)
// 				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "ToAccountCurrencyMismatch",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account3.ID,
// 				"amount":          amount,
// 				"currency":        enums.USD,
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, authorizationTypeBearer, tokenMaker, user1.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(account1, nil)
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account3.ID)).Times(1).Return(account3, nil)
// 				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "InvalidCurrency",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        "XYZ",
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, authorizationTypeBearer, tokenMaker, user1.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
// 				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "NegativeAmount",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          -amount,
// 				"currency":        enums.USD,
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, authorizationTypeBearer, tokenMaker, user1.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
// 				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "GetAccountError",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        enums.USD,
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, authorizationTypeBearer, tokenMaker, user1.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(1).Return(db.Account{}, sql.ErrConnDone)
// 				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "TransferTxError",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        enums.USD,
// 			},
// 			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
// 				addAuthorization(t, request, authorizationTypeBearer, tokenMaker, user1.Username, time.Minute)
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(account1, nil)
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.ID)).Times(1).Return(account2, nil)
// 				store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(1).Return(db.TransferTxResult{}, sql.ErrTxDone)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)

// 			server := newTestServer(t, store)
// 			recorder := httptest.NewRecorder()

// 			// Marshal body data to JSON
// 			data, err := json.Marshal(tc.body)
// 			require.NoError(t, err)

// 			url := "/transfers"
// 			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			require.NoError(t, err)

// 			tc.setupAuth(t, request, server.TokenMaker)
// 			server.Router.ServeHTTP(recorder, request)
// 			tc.checkResponse(recorder)
// 		})
// 	}
// }
