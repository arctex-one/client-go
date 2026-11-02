// Code generated from the maintained OpenAPI contract; DO NOT EDIT.
package generated

import (
	"context"
	"encoding/json"
	"github.com/arctex-one/client-go/transport"
	"net/http"
	"net/url"
	"strings"
)

const Contract = "{\n  \"openapi\": \"3.1.0\",\n  \"info\": {\n    \"title\": \"Arctex API\",\n    \"version\": \"0.5.0\",\n    \"summary\": \"Digital asset payment processing for supported assets and blockchain networks.\",\n    \"description\": \"Arctex One provides digital asset payment processing through a single API. Use the Arctex API to work with financial accounts, create receiving addresses, request quotes, submit transfers and receive webhook notifications for supported assets and blockchain networks.\\n\\n**Integration preview — draft contract v0.5.0.** All 17 HTTP operations and three webhook types are implemented. Recorded acceptance includes isolated HTTP and financial recovery, credential/signing-key rotation and funded test-network bridge/payout recovery after restart. Your application still requires integration acceptance, and production activation requires operational approval. This preview does not claim a production API deployment or service-level agreement.\\n\\n[Download integration guide (Markdown)](/guide.md) · [Download JSON](/openapi.json) · [Download YAML](/openapi.yaml)\\n\\nYour bearer credential determines your tenant, initiating actor and permitted financial accounts. Amounts and versions are strings; retain exact decimal text and the complete network-specific asset identity. Financial accounts are separate from physical wallets.\\n\\nPersist the original Idempotency-Key, request and correlation before each POST. Replay that same identity after uncertainty. A 202 accepts durable transfer processing; it does not promise a committed financial hold, immediate execution or settlement. Accepted instructions and economics remain frozen through recovery.\\n\\nOnly committed financial effects provide financial evidence. A transfer state, observed deposit or confirmed blockchain transaction cannot authorize a customer posting. Authenticate webhook bytes before decoding, store a durable inbox, deduplicate effect_ref with your financial posting and use recovery reads to reconcile missed events.\\n\\nSupport is installation-specific. Ordinary crypto quotation supports FIX_SOURCE with either fee treatment and FIX_RECEIVED with ADD_TO_SOURCE. Ordinary FIX_RECEIVED with DEDUCT_FROM_DESTINATION returns 422 UNSUPPORTED_QUOTE_COMBINATION; configured bridge quotation supports all four combinations. No universal network coverage or instant settlement is promised.\\n\\nUse the endpoint and credentials issued during onboarding. The example origin below is deliberately non-routable; request execution is disabled in Arctex Docs. See the integration guide for current cursor/delivery policies, error handling and production acceptance requirements.\"\n  },\n  \"servers\": [\n    {\n      \"url\": \"https://api.arctex.invalid\",\n      \"description\": \"Non-routable documentation placeholder. Replace with your installation-issued endpoint; production API availability is not claimed by this preview.\"\n    }\n  ],\n  \"security\": [\n    {\n      \"TenantBearer\": []\n    }\n  ],\n  \"x-status\": \"draft-implemented\",\n  \"x-source-revision\": \"8750a452379eaf40cb1c16eee32fb874a2b726ce\",\n  \"tags\": [\n    {\n      \"name\": \"Assets\",\n      \"description\": \"Authoritative enabled crypto assets; static eligibility is not execution readiness.\"\n    },\n    {\n      \"name\": \"Accounts\",\n      \"description\": \"Arctex-owned financial accounts and committed financial history.\"\n    },\n    {\n      \"name\": \"Receiving addresses\",\n      \"description\": \"Tenant receiving routes with immutable attribution history; physical execution remains private.\"\n    },\n    {\n      \"name\": \"Financial effects\",\n      \"description\": \"Recover immutable, tenant-visible committed Accounting effects; resource status is not financial evidence.\"\n    },\n    {\n      \"name\": \"Deposits\",\n      \"description\": \"Existing deposit processing and committed credit evidence.\"\n    },\n    {\n      \"name\": \"Transfers\",\n      \"description\": \"Crypto quotes and quotation-backed Processing transfers.\"\n    },\n    {\n      \"name\": \"Bridge routes\",\n      \"description\": \"Existing public bridge route choices; protocol execution stays private.\"\n    },\n    {\n      \"name\": \"Notifications\",\n      \"description\": \"Outbound transfer updates, deposit updates and committed financial effects. Transfer and deposit updates describe processing state; only financial effects provide committed Accounting evidence.\"\n    }\n  ],\n  \"paths\": {\n    \"/v1/assets\": {\n      \"get\": {\n        \"operationId\": \"listAssets\",\n        \"tags\": [\n          \"Assets\"\n        ],\n        \"summary\": \"List enabled network-specific crypto assets\",\n        \"description\": \"Read the active asset catalog for your installation. Availability is limited to supported digital assets and configured networks; catalog enablement alone does not guarantee current execution readiness.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/treasury/v1/currency_catalog.proto: CurrencyCatalogService.ListCurrencyAssets\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Successful tenant-scoped response.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/AssetList\"\n                },\n                \"example\": {\n                  \"catalog_revision\": \"illustrative-revision-01\",\n                  \"items\": [\n                    {\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      },\n                      \"accounting_scale\": 6,\n                      \"display_scale\": 2\n                    }\n                  ]\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          }\n        },\n        \"x-implementation\": \"api/processing/internal/publicapi/server.go: assets; standalone registry reuses the authoritative crypto asset catalog.\"\n      }\n    },\n    \"/v1/accounts\": {\n      \"get\": {\n        \"operationId\": \"listAccounts\",\n        \"tags\": [\n          \"Accounts\"\n        ],\n        \"summary\": \"List authorized financial accounts\",\n        \"description\": \"The boundary supplies organization, viewer, and visibility scope from the authenticated principal. Returns Accounting accounts, not physical wallets.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/ledger/v1/ledger.proto: LedgerService.ListAccounts\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Successful tenant-scoped response.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/AccountList\"\n                },\n                \"example\": {\n                  \"items\": [\n                    {\n                      \"account_ref\": \"acct_example_01\",\n                      \"account_code\": \"tenant-crypto\",\n                      \"status\": \"ACTIVE\",\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      }\n                    }\n                  ]\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          }\n        },\n        \"x-implementation\": \"api/processing/internal/publicapi/server.go: accounts; api/processing/internal/standalone/accounts.go uses MC Accounting with tenant/account/network bindings.\"\n      }\n    },\n    \"/v1/accounts/{account_ref}/balance\": {\n      \"get\": {\n        \"operationId\": \"getBalance\",\n        \"tags\": [\n          \"Accounts\"\n        ],\n        \"summary\": \"Read an Accounting balance\",\n        \"description\": \"Read an Accounting balance. Independently paged effects and this read do not share an Accounting snapshot. Current version and event-time as_of cannot be substituted for a commit cutoff. Follow /guide.md#consistent-reconciliation-cutoff before declaring a balance/effect reconciliation complete.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/ledger/v1/ledger.proto: LedgerService.GetBalance\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Successful tenant-scoped response.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/Balance\"\n                },\n                \"example\": {\n                  \"account_ref\": \"acct_example_01\",\n                  \"balance\": {\n                    \"amount\": \"1000.00\",\n                    \"currency\": \"USDC\"\n                  },\n                  \"version\": \"3\",\n                  \"last_updated\": \"2026-10-23T10:02:00Z\",\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  }\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"account_ref\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"description\": \"Opaque account_ref; the authenticated tenant must own or be authorized to view the resource.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          },\n          {\n            \"name\": \"as_of\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"description\": \"Optional inclusive event-time cutoff, normalized to UTC. Historical results omit materialized version and last_updated. This is an event-time cutoff, not a frozen commit snapshot; a later backdated entry may affect another read at the same as_of.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Timestamp\"\n            }\n          }\n        ],\n        \"x-implementation\": \"api/processing/internal/publicapi/server.go: balance; standalone/accounts.go reads the authorized network-bound Accounting balance.\"\n      }\n    },\n    \"/v1/accounts/{account_ref}/statement-entries\": {\n      \"get\": {\n        \"operationId\": \"listStatementEntries\",\n        \"tags\": [\n          \"Accounts\"\n        ],\n        \"summary\": \"List committed account statement entries\",\n        \"description\": \"Returns committed account-scoped financial history. This owner uses limit/offset, not a cursor.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/ledger/v1/ledger.proto: LedgerService.ListStatementEntries\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Successful tenant-scoped response.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/StatementList\"\n                },\n                \"example\": {\n                  \"items\": [\n                    {\n                      \"entry_ref\": \"journal_example_01\",\n                      \"entry_type\": \"PAYMENT\",\n                      \"account_amount\": {\n                        \"amount\": \"-126.00\",\n                        \"currency\": \"USDC\"\n                      },\n                      \"description\": \"Illustrative transfer debit and Arctex fee\",\n                      \"event_time\": \"2026-10-23T10:02:00Z\",\n                      \"version\": \"1\",\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      },\n                      \"effect_refs\": [\n                        \"journal_example_01:line_example_01\",\n                        \"journal_example_01:line_example_02\"\n                      ]\n                    }\n                  ]\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"account_ref\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"description\": \"Opaque account_ref; the authenticated tenant must own or be authorized to view the resource.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          },\n          {\n            \"$ref\": \"#/components/parameters/Limit\"\n          },\n          {\n            \"name\": \"offset\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"description\": \"Number of statement entries to skip, preserving existing owner semantics.\",\n            \"schema\": {\n              \"type\": \"integer\",\n              \"minimum\": 0,\n              \"default\": 0\n            }\n          }\n        ],\n        \"x-implementation\": \"api/processing/internal/publicapi/financial.go: statement projection from committed Accounting membership; immutable effect/line/origin identities and account-scoped visibility.\"\n      }\n    },\n    \"/v1/accounts/{account_ref}/statement-entries/{entry_ref}\": {\n      \"get\": {\n        \"operationId\": \"getStatementEntry\",\n        \"tags\": [\n          \"Accounts\"\n        ],\n        \"summary\": \"Read a committed account statement entry\",\n        \"description\": \"Apply existing visibility filtering. Do not expose other tenants, balancing accounts, or a general journal-write interface.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/ledger/v1/ledger.proto: LedgerService.GetStatementEntry\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Successful tenant-scoped response.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/StatementEntry\"\n                },\n                \"example\": {\n                  \"entry_ref\": \"journal_example_01\",\n                  \"entry_type\": \"PAYMENT\",\n                  \"account_amount\": {\n                    \"amount\": \"-126.00\",\n                    \"currency\": \"USDC\"\n                  },\n                  \"description\": \"Illustrative transfer debit and Arctex fee\",\n                  \"event_time\": \"2026-10-23T10:02:00Z\",\n                  \"version\": \"1\",\n                  \"lines\": [\n                    {\n                      \"account_ref\": \"acct_example_01\",\n                      \"money\": {\n                        \"amount\": \"-125.50\",\n                        \"currency\": \"USDC\"\n                      },\n                      \"line_type\": \"MAIN\",\n                      \"posting_line_ref\": \"line_example_01\",\n                      \"effect_ref\": \"journal_example_01:line_example_01\",\n                      \"origin\": {\n                        \"type\": \"TRANSFER\",\n                        \"transfer_ref\": \"transfer_example_01\",\n                        \"client_reference\": \"merchant-withdrawal-2026-001\"\n                      },\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      }\n                    },\n                    {\n                      \"account_ref\": \"acct_example_01\",\n                      \"money\": {\n                        \"amount\": \"-0.50\",\n                        \"currency\": \"USDC\"\n                      },\n                      \"line_type\": \"FEE\",\n                      \"posting_line_ref\": \"line_example_02\",\n                      \"effect_ref\": \"journal_example_01:line_example_02\",\n                      \"origin\": {\n                        \"type\": \"TRANSFER\",\n                        \"transfer_ref\": \"transfer_example_01\",\n                        \"client_reference\": \"merchant-withdrawal-2026-001\"\n                      },\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      }\n                    }\n                  ],\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  },\n                  \"effect_refs\": [\n                    \"journal_example_01:line_example_01\",\n                    \"journal_example_01:line_example_02\"\n                  ]\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"account_ref\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"description\": \"Opaque account_ref; the authenticated tenant must own or be authorized to view the resource.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          },\n          {\n            \"name\": \"entry_ref\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"description\": \"Opaque entry_ref; the authenticated tenant must own or be authorized to view the resource.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          }\n        ],\n        \"x-implementation\": \"api/processing/internal/publicapi/financial.go: statement detail projection from the same committed Accounting receipts as financial-effect reads.\"\n      }\n    },\n    \"/v1/receiving-addresses\": {\n      \"get\": {\n        \"operationId\": \"listReceivingAddresses\",\n        \"tags\": [\n          \"Receiving addresses\"\n        ],\n        \"summary\": \"List tenant receiving-address routes\",\n        \"description\": \"List tenant receiving-address routes.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/inbound/routes/v1/routes.proto: RouteService.ListRoutes\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Successful tenant-scoped response.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/ReceivingAddressList\"\n                },\n                \"example\": {\n                  \"items\": [\n                    {\n                      \"receiving_address_ref\": \"route_example_01\",\n                      \"account_ref\": \"acct_example_01\",\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      },\n                      \"address\": \"0x3333333333333333333333333333333333333333\",\n                      \"status\": \"ACTIVE\",\n                      \"version\": \"1\",\n                      \"created_at\": \"2026-10-23T10:00:00Z\"\n                    }\n                  ]\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"account_ref\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"description\": \"Optional authorized receiving financial account filter.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          },\n          {\n            \"$ref\": \"#/components/parameters/Cursor\"\n          },\n          {\n            \"$ref\": \"#/components/parameters/Limit\"\n          }\n        ],\n        \"x-implementation\": \"api/processing/internal/publicapi/server.go: receiving-address resource projection; receiving.go retains the original provisioning/route owner.\"\n      },\n      \"post\": {\n        \"operationId\": \"createReceivingAddress\",\n        \"tags\": [\n          \"Receiving addresses\"\n        ],\n        \"summary\": \"Provision a receiving address and attribution route\",\n        \"description\": \"Return only after durable provisioning and route attribution are established. Reuse the existing coordinator. On timeout/lost response, replay the exact request/key; no new provisioning identity. No physical pool/key selection is accepted.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/edge/product/internal/walletprovisioning + RouteService.CreateRoute\",\n        \"responses\": {\n          \"201\": {\n            \"description\": \"Receiving address and immutable attribution route durably established.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/ReceivingAddress\"\n                },\n                \"example\": {\n                  \"receiving_address_ref\": \"route_example_01\",\n                  \"account_ref\": \"acct_example_01\",\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  },\n                  \"address\": \"0x3333333333333333333333333333333333333333\",\n                  \"status\": \"ACTIVE\",\n                  \"version\": \"1\",\n                  \"created_at\": \"2026-10-23T10:00:00Z\"\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          },\n          \"409\": {\n            \"$ref\": \"#/components/responses/Conflict\"\n          },\n          \"422\": {\n            \"$ref\": \"#/components/responses/InvalidInput\"\n          },\n          \"410\": {\n            \"$ref\": \"#/components/responses/IdempotencyExpired\"\n          }\n        },\n        \"parameters\": [\n          {\n            \"$ref\": \"#/components/parameters/IdempotencyKey\"\n          }\n        ],\n        \"requestBody\": {\n          \"required\": true,\n          \"content\": {\n            \"application/json\": {\n              \"schema\": {\n                \"$ref\": \"#/components/schemas/CreateReceivingAddress\"\n              }\n            }\n          }\n        },\n        \"x-implementation\": \"api/processing/internal/publicapi/receiving.go: durable original provisioning and frozen asset/target replay; standalone/public_boundary.go adapts the retained owner.\"\n      }\n    },\n    \"/v1/receiving-addresses/{receiving_address_ref}\": {\n      \"get\": {\n        \"operationId\": \"getReceivingAddress\",\n        \"tags\": [\n          \"Receiving addresses\"\n        ],\n        \"summary\": \"Read a receiving address and its attribution\",\n        \"description\": \"Read a receiving address and its attribution.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/inbound/routes/v1/routes.proto: RouteService.GetRoute\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Successful tenant-scoped response.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/ReceivingAddress\"\n                },\n                \"example\": {\n                  \"receiving_address_ref\": \"route_example_01\",\n                  \"account_ref\": \"acct_example_01\",\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  },\n                  \"address\": \"0x3333333333333333333333333333333333333333\",\n                  \"status\": \"ACTIVE\",\n                  \"version\": \"1\",\n                  \"created_at\": \"2026-10-23T10:00:00Z\"\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"receiving_address_ref\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"description\": \"Opaque receiving_address_ref; the authenticated tenant must own or be authorized to view the resource.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          }\n        ],\n        \"x-implementation\": \"api/processing/internal/publicapi/server.go: authorized receiving-address read from the retained route/provisioning projection.\"\n      }\n    },\n    \"/v1/deposits\": {\n      \"get\": {\n        \"operationId\": \"listDepositRevisions\",\n        \"tags\": [\n          \"Deposits\"\n        ],\n        \"summary\": \"Discover deposits and recover later revisions\",\n        \"description\": \"Recover missed initial observations and later updates, including delayed credit. Omit cursor to read retained history. Continue using next_cursor, including after has_more=false or an empty page. Persist a returned cursor only after durably handling every item. Tokens bind the authenticated tenant, this surface, and unchanged filters; malformed/mismatched tokens return 400, expired or unrecoverable checkpoints return 410 and require reconciliation. Never silently reset. Cursor order is durable projection commit order, not created_at/event_time, so late commits and state updates remain discoverable. Projection recovery must not advance past unprojected committed source changes. Current recovery cursors expire after 24 hours; retained revisions/effects have no automatic TTL. Production retention, archival and replay agreements remain release gates; see /guide.md#installed-retention-and-delivery-policy. Each item is an immutable retained Deposit revision. The same deposit_ref can recur at a later version, including a late credit; deduplicate/compare by deposit_ref and version. GET by reference returns the current revision. Version need not be contiguous. No generic event API or new deposit lifecycle is introduced. Before route attribution, visibility derives from retained tenant ownership of intake/custody evidence. An unauthorized account/address filter must be rejected or hidden under the standard visibility policy; it cannot grant access to unresolved deposits.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/inbound/crypto_deposits/internal/storage/model.go: InboundDeposit.Version and intake snapshot; repository.go: DepositStore; api/proto/inbound/routes/v1/routes.proto: route/history\",\n        \"parameters\": [\n          {\n            \"$ref\": \"#/components/parameters/RecoveryCursor\"\n          },\n          {\n            \"$ref\": \"#/components/parameters/Limit\"\n          },\n          {\n            \"name\": \"account_ref\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"description\": \"Filter by resolved authorized financial account. Unresolved deposits are excluded; omit this filter for complete tenant discovery.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          },\n          {\n            \"name\": \"receiving_address_ref\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"description\": \"Filter by immutable recognized receiving reference, not the latest route.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          }\n        ],\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Authorized recovery projection.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/DepositList\"\n                },\n                \"example\": {\n                  \"items\": [\n                    {\n                      \"deposit_ref\": \"deposit_example_01\",\n                      \"inbound_ref\": \"inbound_example_01\",\n                      \"account_ref\": \"acct_example_01\",\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      },\n                      \"amount\": \"125.50\",\n                      \"state\": \"AML_PENDING\",\n                      \"created_at\": \"2026-10-23T10:00:00Z\",\n                      \"updated_at\": \"2026-10-23T10:00:05Z\",\n                      \"receiving\": {\n                        \"receiving_address_ref\": \"route_example_01\",\n                        \"route_version\": \"1\",\n                        \"address\": \"0x3333333333333333333333333333333333333333\",\n                        \"captured_at\": \"2026-10-23T10:00:00Z\"\n                      },\n                      \"version\": \"2\"\n                    },\n                    {\n                      \"deposit_ref\": \"deposit_example_01\",\n                      \"inbound_ref\": \"inbound_example_01\",\n                      \"account_ref\": \"acct_example_01\",\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      },\n                      \"amount\": \"125.50\",\n                      \"state\": \"COMPLETED\",\n                      \"created_at\": \"2026-10-23T10:00:00Z\",\n                      \"updated_at\": \"2026-10-24T11:00:00Z\",\n                      \"receiving\": {\n                        \"receiving_address_ref\": \"route_example_01\",\n                        \"route_version\": \"1\",\n                        \"address\": \"0x3333333333333333333333333333333333333333\",\n                        \"captured_at\": \"2026-10-23T10:00:00Z\"\n                      },\n                      \"version\": \"7\",\n                      \"credit_journal_entry_ref\": \"journal_deposit_example_01\"\n                    }\n                  ],\n                  \"next_cursor\": \"opaque-deposit-checkpoint-02\",\n                  \"has_more\": false\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          },\n          \"410\": {\n            \"$ref\": \"#/components/responses/RecoveryCursorExpired\"\n          }\n        },\n        \"x-implementation\": \"api/processing/internal/publicapi/deposits.go and store.go: immutable owner-revision projection with scoped durable cursors and replay after restart. Historical pre-checkpoint gaps require explicit adoption.\"\n      }\n    },\n    \"/v1/deposits/{deposit_ref}\": {\n      \"get\": {\n        \"operationId\": \"getDeposit\",\n        \"tags\": [\n          \"Deposits\"\n        ],\n        \"summary\": \"Read deposit processing and credit evidence\",\n        \"description\": \"Read an existing durable deposit identity. Raw chain observations and bootstrap balances do not imply a customer credit. Route attribution must be checked within the tenant before returning data. The same authorized principal can resolve this reference before a successful POST response, recovery item, or webhook that introduces it is published. Transient projection/dependency lag returns 503 and is retryable; 404 means absent or outside authorized visibility, not expected publication lag. Cross-surface delivery order remains unspecified.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/inbound/crypto_deposits/v1/crypto_deposits.proto: InboundDepositProcessorService.GetInboundDeposit\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Successful tenant-scoped response.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/Deposit\"\n                },\n                \"example\": {\n                  \"deposit_ref\": \"deposit_example_01\",\n                  \"inbound_ref\": \"inbound_example_01\",\n                  \"account_ref\": \"acct_example_01\",\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  },\n                  \"amount\": \"125.50\",\n                  \"state\": \"AML_PENDING\",\n                  \"created_at\": \"2026-10-23T10:00:00Z\",\n                  \"updated_at\": \"2026-10-23T10:00:05Z\",\n                  \"receiving\": {\n                    \"receiving_address_ref\": \"route_example_01\",\n                    \"route_version\": \"1\",\n                    \"address\": \"0x3333333333333333333333333333333333333333\",\n                    \"captured_at\": \"2026-10-23T10:00:00Z\"\n                  },\n                  \"version\": \"2\"\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"deposit_ref\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"description\": \"Opaque deposit_ref; the authenticated tenant must own or be authorized to view the resource.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          }\n        ],\n        \"x-implementation\": \"api/processing/internal/publicapi/deposits.go: current retained owner revision, immutable receiving linkage and persisted chain evidence; missing legacy history is not fabricated.\"\n      }\n    },\n    \"/v1/quotes\": {\n      \"post\": {\n        \"operationId\": \"createQuote\",\n        \"tags\": [\n          \"Transfers\"\n        ],\n        \"summary\": \"Quote a crypto transfer\",\n        \"description\": \"Create a persisted crypto-only quote; no financial/physical reservation or execution admission is implied. Addresses in examples are illustrative, not deployed test fixtures. A quote is not an instruction to post a customer ledger entry. Ordinary crypto supports FIX_SOURCE with either fee treatment and FIX_RECEIVED with ADD_TO_SOURCE; ordinary exact-receive/deduction returns 422 UNSUPPORTED_QUOTE_COMBINATION. Bridge quotation supports all four combinations subject to installed route policy.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/payments/quotation/v2/quotation.proto: QuotationService.QuotePayment\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Successful tenant-scoped response.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/Quote\"\n                },\n                \"example\": {\n                  \"quote_ref\": \"quote_example_01\",\n                  \"state\": \"EXECUTABLE\",\n                  \"principal\": {\n                    \"amount\": \"125.50\",\n                    \"currency\": \"USDC\"\n                  },\n                  \"recipient_amount\": {\n                    \"amount\": \"125.50\",\n                    \"currency\": \"USDC\"\n                  },\n                  \"payer_total_debit\": {\n                    \"amount\": \"126.00\",\n                    \"currency\": \"USDC\"\n                  },\n                  \"fees\": [\n                    {\n                      \"code\": \"mc_processing\",\n                      \"amount\": {\n                        \"amount\": \"0.50\",\n                        \"currency\": \"USDC\"\n                      },\n                      \"description\": \"Illustrative Arctex processing fee\",\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      }\n                    }\n                  ],\n                  \"fee_treatment\": \"ADD_TO_SOURCE\",\n                  \"settlement_mode\": \"FIX_SOURCE\",\n                  \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n                  \"priced_at\": \"2026-10-23T10:00:00Z\",\n                  \"expires_at\": \"2026-10-23T10:01:00Z\",\n                  \"instructions\": {\n                    \"source_account_ref\": \"acct_example_01\",\n                    \"source_asset\": {\n                      \"network\": \"ethereum_sepolia\",\n                      \"currency\": \"USDC\",\n                      \"kind\": \"TOKEN_CONTRACT\",\n                      \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                    },\n                    \"destination\": {\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      },\n                      \"address\": \"0x2222222222222222222222222222222222222222\"\n                    }\n                  }\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          },\n          \"409\": {\n            \"$ref\": \"#/components/responses/Conflict\"\n          },\n          \"422\": {\n            \"$ref\": \"#/components/responses/InvalidInput\"\n          },\n          \"410\": {\n            \"$ref\": \"#/components/responses/IdempotencyExpired\"\n          }\n        },\n        \"parameters\": [\n          {\n            \"$ref\": \"#/components/parameters/IdempotencyKey\"\n          }\n        ],\n        \"requestBody\": {\n          \"required\": true,\n          \"content\": {\n            \"application/json\": {\n              \"schema\": {\n                \"$ref\": \"#/components/schemas/QuoteRequest\"\n              },\n              \"example\": {\n                \"source_account_ref\": \"acct_example_01\",\n                \"source_asset\": {\n                  \"network\": \"ethereum_sepolia\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"destination\": {\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  },\n                  \"address\": \"0x2222222222222222222222222222222222222222\"\n                },\n                \"amount\": \"125.50\",\n                \"settlement_mode\": \"FIX_SOURCE\",\n                \"fee_treatment\": \"ADD_TO_SOURCE\",\n                \"client_reference\": \"merchant-withdrawal-2026-001\"\n              }\n            }\n          }\n        },\n        \"x-implementation\": \"api/processing/internal/standalone/public_boundary.go: original quotation/fee/bridge calculators with publicapi/server.go durable transport identity. Ordinary exact-receive/deduction is rejected; bridge supports all four modes.\"\n      }\n    },\n    \"/v1/transfers\": {\n      \"post\": {\n        \"operationId\": \"submitTransfer\",\n        \"tags\": [\n          \"Transfers\"\n        ],\n        \"summary\": \"Admit execution of an existing quote\",\n        \"description\": \"Persist the caller key and request before sending. 202 acknowledges the existing durable Processing aggregate, not blockchain success or committed Accounting settlement. Retry the same POST after response loss. An admitted transfer retains its quote snapshot after expiry; acceptance does not promise immediate liquidity. A public quote represents one existing intent and can own at most one admitted transfer. While valid, semantically identical submissions under different keys converge on the same transfer, including concurrent requests; the original client_reference and accepted snapshot win. Incompatible claims return 409 QUOTE_CLAIM_CONFLICT. Same-key conflicting input remains IDEMPOTENCY_CONFLICT. Same-key admitted replay is resolved before expiry; a previously unseen different-key request after quote expiry may return QUOTE_EXPIRED, so recover by GET/list filtered by quote_ref rather than changing keys. 202 establishes durable acceptance/scheduled recovery, not proof that an Accounting hold has committed. Reservation/spending checks occur in the existing financial workflow before dispatch; failures after 202 are reported on the same transfer with attention, not a retroactive HTTP rejection. See the plan for financial outcomes and recovery. Every key acknowledged with 202, including a converged duplicate key, must durably resolve to the original transfer within the published idempotency window. A source business claim alone does not prove this HTTP key-to-result retention; it is an explicit boundary acceptance gate.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/payments/orchestration/v2/orchestration.proto: PaymentOrchestratorService.ExecutePayment; orchestrationv2/psvc/business_claims.go and execute_async.go; prepo/indexes.go\",\n        \"responses\": {\n          \"202\": {\n            \"description\": \"Existing durable Processing transfer admitted. No broadcast or committed financial settlement is implied.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/Transfer\"\n                },\n                \"example\": {\n                  \"transfer_ref\": \"transfer_example_01\",\n                  \"quote_ref\": \"quote_example_01\",\n                  \"client_reference\": \"merchant-withdrawal-2026-001\",\n                  \"state\": \"CREATED\",\n                  \"version\": \"1\",\n                  \"quote\": {\n                    \"quote_ref\": \"quote_example_01\",\n                    \"state\": \"EXECUTABLE\",\n                    \"principal\": {\n                      \"amount\": \"125.50\",\n                      \"currency\": \"USDC\"\n                    },\n                    \"recipient_amount\": {\n                      \"amount\": \"125.50\",\n                      \"currency\": \"USDC\"\n                    },\n                    \"payer_total_debit\": {\n                      \"amount\": \"126.00\",\n                      \"currency\": \"USDC\"\n                    },\n                    \"fees\": [\n                      {\n                        \"code\": \"mc_processing\",\n                        \"amount\": {\n                          \"amount\": \"0.50\",\n                          \"currency\": \"USDC\"\n                        },\n                        \"description\": \"Illustrative Arctex processing fee\",\n                        \"asset\": {\n                          \"network\": \"ethereum_sepolia\",\n                          \"currency\": \"USDC\",\n                          \"kind\": \"TOKEN_CONTRACT\",\n                          \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                        }\n                      }\n                    ],\n                    \"fee_treatment\": \"ADD_TO_SOURCE\",\n                    \"settlement_mode\": \"FIX_SOURCE\",\n                    \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n                    \"priced_at\": \"2026-10-23T10:00:00Z\",\n                    \"expires_at\": \"2026-10-23T10:01:00Z\",\n                    \"instructions\": {\n                      \"source_account_ref\": \"acct_example_01\",\n                      \"source_asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      },\n                      \"destination\": {\n                        \"asset\": {\n                          \"network\": \"ethereum_sepolia\",\n                          \"currency\": \"USDC\",\n                          \"kind\": \"TOKEN_CONTRACT\",\n                          \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                        },\n                        \"address\": \"0x2222222222222222222222222222222222222222\"\n                      }\n                    }\n                  },\n                  \"executions\": [],\n                  \"bridges\": [],\n                  \"created_at\": \"2026-10-23T10:00:02Z\",\n                  \"updated_at\": \"2026-10-23T10:00:02Z\"\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          },\n          \"409\": {\n            \"$ref\": \"#/components/responses/Conflict\"\n          },\n          \"422\": {\n            \"$ref\": \"#/components/responses/InvalidInput\"\n          },\n          \"410\": {\n            \"$ref\": \"#/components/responses/IdempotencyExpired\"\n          }\n        },\n        \"parameters\": [\n          {\n            \"$ref\": \"#/components/parameters/IdempotencyKey\"\n          }\n        ],\n        \"requestBody\": {\n          \"required\": true,\n          \"content\": {\n            \"application/json\": {\n              \"schema\": {\n                \"$ref\": \"#/components/schemas/SubmitTransfer\"\n              },\n              \"example\": {\n                \"quote_ref\": \"quote_example_01\",\n                \"client_reference\": \"merchant-withdrawal-2026-001\"\n              }\n            }\n          }\n        },\n        \"x-implementation\": \"api/processing/internal/publicapi/server.go and recovery.go: durable first-write request/result and quote claim, original actor and accepted snapshot; standalone/public_boundary.go delegates admission to the original Processing owner. Recovery reads admitted claims without submitting again.\"\n      },\n      \"get\": {\n        \"operationId\": \"listTransfers\",\n        \"tags\": [\n          \"Transfers\"\n        ],\n        \"summary\": \"List authorized transfers\",\n        \"description\": \"List authorized transfers.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/payments/orchestration/v2/orchestration.proto: PaymentOrchestratorService.ListPayments\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Successful tenant-scoped response.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/TransferList\"\n                },\n                \"example\": {\n                  \"items\": [\n                    {\n                      \"transfer_ref\": \"transfer_example_01\",\n                      \"quote_ref\": \"quote_example_01\",\n                      \"client_reference\": \"merchant-withdrawal-2026-001\",\n                      \"state\": \"CREATED\",\n                      \"version\": \"1\",\n                      \"quote\": {\n                        \"quote_ref\": \"quote_example_01\",\n                        \"state\": \"EXECUTABLE\",\n                        \"principal\": {\n                          \"amount\": \"125.50\",\n                          \"currency\": \"USDC\"\n                        },\n                        \"recipient_amount\": {\n                          \"amount\": \"125.50\",\n                          \"currency\": \"USDC\"\n                        },\n                        \"payer_total_debit\": {\n                          \"amount\": \"126.00\",\n                          \"currency\": \"USDC\"\n                        },\n                        \"fees\": [\n                          {\n                            \"code\": \"mc_processing\",\n                            \"amount\": {\n                              \"amount\": \"0.50\",\n                              \"currency\": \"USDC\"\n                            },\n                            \"description\": \"Illustrative Arctex processing fee\",\n                            \"asset\": {\n                              \"network\": \"ethereum_sepolia\",\n                              \"currency\": \"USDC\",\n                              \"kind\": \"TOKEN_CONTRACT\",\n                              \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                            }\n                          }\n                        ],\n                        \"fee_treatment\": \"ADD_TO_SOURCE\",\n                        \"settlement_mode\": \"FIX_SOURCE\",\n                        \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n                        \"priced_at\": \"2026-10-23T10:00:00Z\",\n                        \"expires_at\": \"2026-10-23T10:01:00Z\",\n                        \"instructions\": {\n                          \"source_account_ref\": \"acct_example_01\",\n                          \"source_asset\": {\n                            \"network\": \"ethereum_sepolia\",\n                            \"currency\": \"USDC\",\n                            \"kind\": \"TOKEN_CONTRACT\",\n                            \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                          },\n                          \"destination\": {\n                            \"asset\": {\n                              \"network\": \"ethereum_sepolia\",\n                              \"currency\": \"USDC\",\n                              \"kind\": \"TOKEN_CONTRACT\",\n                              \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                            },\n                            \"address\": \"0x2222222222222222222222222222222222222222\"\n                          }\n                        }\n                      },\n                      \"executions\": [],\n                      \"bridges\": [],\n                      \"created_at\": \"2026-10-23T10:00:02Z\",\n                      \"updated_at\": \"2026-10-23T10:00:02Z\"\n                    }\n                  ]\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          }\n        },\n        \"parameters\": [\n          {\n            \"$ref\": \"#/components/parameters/Cursor\"\n          },\n          {\n            \"$ref\": \"#/components/parameters/Limit\"\n          },\n          {\n            \"name\": \"quote_ref\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"description\": \"Filter by original quotation identity.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          },\n          {\n            \"name\": \"transfer_ref\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"description\": \"Filter by existing transfer identity.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          },\n          {\n            \"name\": \"state\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"description\": \"Existing coordinator state filter.\",\n            \"schema\": {\n              \"type\": \"string\",\n              \"enum\": [\n                \"CREATED\",\n                \"EXECUTING\",\n                \"NEEDS_ATTENTION\",\n                \"SETTLED\",\n                \"FAILED\"\n              ]\n            }\n          }\n        ],\n        \"x-implementation\": \"api/processing/internal/publicapi/server.go and recovery.go: tenant/account-scoped retained transfer projection with quote/client filters and original admission recovery.\"\n      }\n    },\n    \"/v1/transfers/{transfer_ref}\": {\n      \"get\": {\n        \"operationId\": \"getTransfer\",\n        \"tags\": [\n          \"Transfers\"\n        ],\n        \"summary\": \"Read transfer state and retained quote terms\",\n        \"description\": \"Always supply the tenant organization privately, even though the underlying internal service permits a global lookup. State and chain evidence are separate from committed Accounting effects. The same authorized principal can resolve this reference before a successful POST response, recovery item, or webhook that introduces it is published. Transient projection/dependency lag returns 503 and is retryable; 404 means absent or outside authorized visibility, not expected publication lag. Cross-surface delivery order remains unspecified.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/payments/orchestration/v2/orchestration.proto: PaymentOrchestratorService.GetPayment\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Successful tenant-scoped response.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/Transfer\"\n                },\n                \"example\": {\n                  \"transfer_ref\": \"transfer_example_01\",\n                  \"quote_ref\": \"quote_example_01\",\n                  \"client_reference\": \"merchant-withdrawal-2026-001\",\n                  \"state\": \"CREATED\",\n                  \"version\": \"1\",\n                  \"quote\": {\n                    \"quote_ref\": \"quote_example_01\",\n                    \"state\": \"EXECUTABLE\",\n                    \"principal\": {\n                      \"amount\": \"125.50\",\n                      \"currency\": \"USDC\"\n                    },\n                    \"recipient_amount\": {\n                      \"amount\": \"125.50\",\n                      \"currency\": \"USDC\"\n                    },\n                    \"payer_total_debit\": {\n                      \"amount\": \"126.00\",\n                      \"currency\": \"USDC\"\n                    },\n                    \"fees\": [\n                      {\n                        \"code\": \"mc_processing\",\n                        \"amount\": {\n                          \"amount\": \"0.50\",\n                          \"currency\": \"USDC\"\n                        },\n                        \"description\": \"Illustrative Arctex processing fee\",\n                        \"asset\": {\n                          \"network\": \"ethereum_sepolia\",\n                          \"currency\": \"USDC\",\n                          \"kind\": \"TOKEN_CONTRACT\",\n                          \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                        }\n                      }\n                    ],\n                    \"fee_treatment\": \"ADD_TO_SOURCE\",\n                    \"settlement_mode\": \"FIX_SOURCE\",\n                    \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n                    \"priced_at\": \"2026-10-23T10:00:00Z\",\n                    \"expires_at\": \"2026-10-23T10:01:00Z\",\n                    \"instructions\": {\n                      \"source_account_ref\": \"acct_example_01\",\n                      \"source_asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      },\n                      \"destination\": {\n                        \"asset\": {\n                          \"network\": \"ethereum_sepolia\",\n                          \"currency\": \"USDC\",\n                          \"kind\": \"TOKEN_CONTRACT\",\n                          \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                        },\n                        \"address\": \"0x2222222222222222222222222222222222222222\"\n                      }\n                    }\n                  },\n                  \"executions\": [],\n                  \"bridges\": [],\n                  \"created_at\": \"2026-10-23T10:00:02Z\",\n                  \"updated_at\": \"2026-10-23T10:00:02Z\"\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          }\n        },\n        \"parameters\": [\n          {\n            \"name\": \"transfer_ref\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"description\": \"Opaque transfer_ref; the authenticated tenant must own or be authorized to view the resource.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          }\n        ],\n        \"x-implementation\": \"api/processing/internal/publicapi/server.go and sources.go: retained accepted snapshot and original Processing state under current authorization.\"\n      }\n    },\n    \"/v1/bridge-routes\": {\n      \"get\": {\n        \"operationId\": \"listBridgeRoutes\",\n        \"tags\": [\n          \"Bridge routes\"\n        ],\n        \"summary\": \"List configured crypto bridge routes\",\n        \"description\": \"Routes are selected through quote creation and executed by the existing coordinator. No raw burn/mint, wallet recipient override, or standalone Treasury operation is exposed.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/proto/payments/quotation/v2/quotation.proto: QuotationService.ListBridgeRoutes\",\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Successful tenant-scoped response.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/BridgeRouteList\"\n                },\n                \"example\": {\n                  \"items\": []\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          }\n        },\n        \"x-implementation\": \"api/processing/internal/publicapi/server.go: bridgeRoutes; standalone/public_boundary.go delegates to the original quoter and installation bridge policy.\"\n      }\n    },\n    \"/v1/financial-effects\": {\n      \"get\": {\n        \"operationId\": \"listFinancialEffects\",\n        \"tags\": [\n          \"Financial effects\"\n        ],\n        \"summary\": \"Recover committed financial effects\",\n        \"description\": \"Discover only tenant-visible committed Accounting effects. Omit cursor to read retained history. Continue using next_cursor, including after has_more=false or an empty page. Persist a returned cursor only after durably handling every item. Tokens bind the authenticated tenant, this surface, and unchanged filters; malformed/mismatched tokens return 400, expired or unrecoverable checkpoints return 410 and require reconciliation. Never silently reset. Cursor order is durable projection commit order, not created_at/event_time, so late commits and state updates remain discoverable. Projection recovery must not advance past unprojected committed source changes. Current recovery cursors expire after 24 hours; retained revisions/effects have no automatic TTL. Production retention, archival and replay agreements remain release gates; see /guide.md#installed-retention-and-delivery-policy. Each item is an immutable committed effect. The same effect_ref always has the same financial content, whether read here, by reference, or delivered by webhook. Neither late event_time nor a failed webhook removes it from recovery. has_more=false reflects this projection, not a balance snapshot or source completeness proof. Follow /guide.md#consistent-reconciliation-cutoff.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/pkg/model/ledger_events.go: committed LedgerJournalEvent / LedgerPostingLineEvent and posting context\",\n        \"parameters\": [\n          {\n            \"$ref\": \"#/components/parameters/RecoveryCursor\"\n          },\n          {\n            \"$ref\": \"#/components/parameters/Limit\"\n          },\n          {\n            \"name\": \"account_ref\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"description\": \"Filter by authorized posting account.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          },\n          {\n            \"name\": \"deposit_ref\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"description\": \"Filter DEPOSIT origins; all supplied filters are combined with AND.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          },\n          {\n            \"name\": \"transfer_ref\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"description\": \"Filter TRANSFER origins; cannot match a DEPOSIT origin.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          },\n          {\n            \"name\": \"journal_entry_ref\",\n            \"in\": \"query\",\n            \"required\": false,\n            \"description\": \"Filter by committed journal identity.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          }\n        ],\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Authorized recovery projection.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/FinancialEffectList\"\n                },\n                \"example\": {\n                  \"items\": [\n                    {\n                      \"effect_ref\": \"journal_example_01:line_example_01\",\n                      \"journal_entry_ref\": \"journal_example_01\",\n                      \"posting_line_ref\": \"line_example_01\",\n                      \"account_ref\": \"acct_example_01\",\n                      \"money\": {\n                        \"amount\": \"-125.50\",\n                        \"currency\": \"USDC\"\n                      },\n                      \"committed_at\": \"2026-10-23T10:02:00Z\",\n                      \"origin\": {\n                        \"type\": \"TRANSFER\",\n                        \"transfer_ref\": \"transfer_example_01\",\n                        \"client_reference\": \"merchant-withdrawal-2026-001\"\n                      },\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      },\n                      \"line_type\": \"MAIN\"\n                    },\n                    {\n                      \"effect_ref\": \"journal_example_01:line_example_02\",\n                      \"journal_entry_ref\": \"journal_example_01\",\n                      \"posting_line_ref\": \"line_example_02\",\n                      \"account_ref\": \"acct_example_01\",\n                      \"money\": {\n                        \"amount\": \"-0.50\",\n                        \"currency\": \"USDC\"\n                      },\n                      \"committed_at\": \"2026-10-23T10:02:00Z\",\n                      \"origin\": {\n                        \"type\": \"TRANSFER\",\n                        \"transfer_ref\": \"transfer_example_01\",\n                        \"client_reference\": \"merchant-withdrawal-2026-001\"\n                      },\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      },\n                      \"line_type\": \"FEE\"\n                    },\n                    {\n                      \"effect_ref\": \"journal_deposit_example_01:line_credit_01\",\n                      \"journal_entry_ref\": \"journal_deposit_example_01\",\n                      \"posting_line_ref\": \"line_credit_01\",\n                      \"account_ref\": \"acct_example_01\",\n                      \"money\": {\n                        \"amount\": \"125.50\",\n                        \"currency\": \"USDC\"\n                      },\n                      \"committed_at\": \"2026-10-24T11:00:00Z\",\n                      \"origin\": {\n                        \"type\": \"DEPOSIT\",\n                        \"deposit_ref\": \"deposit_example_01\"\n                      },\n                      \"asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      },\n                      \"line_type\": \"MAIN\"\n                    }\n                  ],\n                  \"next_cursor\": \"opaque-effect-checkpoint-03\",\n                  \"has_more\": false\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          },\n          \"410\": {\n            \"$ref\": \"#/components/responses/RecoveryCursorExpired\"\n          }\n        },\n        \"x-implementation\": \"api/processing/internal/publicapi/financial.go and store.go: replayable immutable committed-effect projection; original posting identity and source membership survive late/backdated commits and projection restart.\"\n      }\n    },\n    \"/v1/financial-effects/{effect_ref}\": {\n      \"get\": {\n        \"operationId\": \"getFinancialEffect\",\n        \"tags\": [\n          \"Financial effects\"\n        ],\n        \"summary\": \"Read a committed financial effect\",\n        \"description\": \"Return exactly the financial object delivered by financial_effect.committed. Reconcile by stable effect_ref and typed origin, never by amount, transaction hash, or resource status. Visibility is enforced on the posting account. The same authorized principal can resolve this reference before a successful POST response, recovery item, or webhook that introduces it is published. Transient projection/dependency lag returns 503 and is retryable; 404 means absent or outside authorized visibility, not expected publication lag. Cross-surface delivery order remains unspecified.\",\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"api/pkg/model/ledger_events.go: committed LedgerJournalEvent / LedgerPostingLineEvent\",\n        \"parameters\": [\n          {\n            \"name\": \"effect_ref\",\n            \"in\": \"path\",\n            \"required\": true,\n            \"description\": \"Stable committed effect identity in the authorized account scope.\",\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            }\n          }\n        ],\n        \"responses\": {\n          \"200\": {\n            \"description\": \"Canonical committed financial effect.\",\n            \"content\": {\n              \"application/json\": {\n                \"schema\": {\n                  \"$ref\": \"#/components/schemas/FinancialEffect\"\n                },\n                \"example\": {\n                  \"effect_ref\": \"journal_example_01:line_example_01\",\n                  \"journal_entry_ref\": \"journal_example_01\",\n                  \"posting_line_ref\": \"line_example_01\",\n                  \"account_ref\": \"acct_example_01\",\n                  \"money\": {\n                    \"amount\": \"-125.50\",\n                    \"currency\": \"USDC\"\n                  },\n                  \"committed_at\": \"2026-10-23T10:02:00Z\",\n                  \"origin\": {\n                    \"type\": \"TRANSFER\",\n                    \"transfer_ref\": \"transfer_example_01\",\n                    \"client_reference\": \"merchant-withdrawal-2026-001\"\n                  },\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  },\n                  \"line_type\": \"MAIN\"\n                }\n              }\n            }\n          },\n          \"400\": {\n            \"$ref\": \"#/components/responses/BadRequest\"\n          },\n          \"401\": {\n            \"$ref\": \"#/components/responses/Unauthorized\"\n          },\n          \"403\": {\n            \"$ref\": \"#/components/responses/Forbidden\"\n          },\n          \"404\": {\n            \"$ref\": \"#/components/responses/NotFound\"\n          },\n          \"429\": {\n            \"$ref\": \"#/components/responses/RateLimited\"\n          },\n          \"503\": {\n            \"$ref\": \"#/components/responses/Unavailable\"\n          }\n        },\n        \"x-implementation\": \"api/processing/internal/publicapi/financial.go: authorized immutable effect projection keyed by the original committed posting-line identity, identical to signed financial notification data.\"\n      }\n    }\n  },\n  \"webhooks\": {\n    \"transferUpdated\": {\n      \"post\": {\n        \"operationId\": \"transferUpdated\",\n        \"tags\": [\n          \"Notifications\"\n        ],\n        \"summary\": \"Transfer updates\",\n        \"description\": \"Resource update only. Never synthesize financial settlement from successful execution. At-least-once delivery may arrive out of order. Verify the HMAC and timestamp before accepting, durably record/deduplicate the event, then return 2xx. Subscription setup is out-of-band. Delivery is at least once and may be out of order. Acknowledge only after durable inbox storage; automatic retries use backoff after timeout/non-2xx, then require administrative replay or read recovery when exhausted. Numeric timeout/backoff/retry horizon and read/inbox retention must be verified and published before production; no numeric delivery SLA is asserted here. Publish only after its resource and its authorized deposit/transfer origins can be read by the same principal. A financial event may precede a resource-update webhook; resolve its origin by GET rather than waiting for notification order.\",\n        \"security\": [],\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"Processing committed Payment resource update\",\n        \"parameters\": [\n          {\n            \"$ref\": \"#/components/parameters/CallbackTimestamp\"\n          },\n          {\n            \"$ref\": \"#/components/parameters/CallbackSignature\"\n          },\n          {\n            \"$ref\": \"#/components/parameters/CallbackAlgorithm\"\n          }\n        ],\n        \"requestBody\": {\n          \"required\": true,\n          \"content\": {\n            \"application/json\": {\n              \"schema\": {\n                \"$ref\": \"#/components/schemas/TransferEvent\"\n              },\n              \"example\": {\n                \"event_id\": \"event_transferUpdated_01\",\n                \"delivery_id\": \"delivery_attempt_01\",\n                \"occurred_at\": \"2026-10-23T10:01:10Z\",\n                \"type\": \"transfer.updated\",\n                \"data\": {\n                  \"transfer_ref\": \"transfer_example_01\",\n                  \"quote_ref\": \"quote_example_01\",\n                  \"client_reference\": \"merchant-withdrawal-2026-001\",\n                  \"state\": \"EXECUTING\",\n                  \"version\": \"2\",\n                  \"quote\": {\n                    \"quote_ref\": \"quote_example_01\",\n                    \"state\": \"EXECUTABLE\",\n                    \"principal\": {\n                      \"amount\": \"125.50\",\n                      \"currency\": \"USDC\"\n                    },\n                    \"recipient_amount\": {\n                      \"amount\": \"125.50\",\n                      \"currency\": \"USDC\"\n                    },\n                    \"payer_total_debit\": {\n                      \"amount\": \"126.00\",\n                      \"currency\": \"USDC\"\n                    },\n                    \"fees\": [\n                      {\n                        \"code\": \"mc_processing\",\n                        \"amount\": {\n                          \"amount\": \"0.50\",\n                          \"currency\": \"USDC\"\n                        },\n                        \"description\": \"Illustrative Arctex processing fee\",\n                        \"asset\": {\n                          \"network\": \"ethereum_sepolia\",\n                          \"currency\": \"USDC\",\n                          \"kind\": \"TOKEN_CONTRACT\",\n                          \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                        }\n                      }\n                    ],\n                    \"fee_treatment\": \"ADD_TO_SOURCE\",\n                    \"settlement_mode\": \"FIX_SOURCE\",\n                    \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n                    \"priced_at\": \"2026-10-23T10:00:00Z\",\n                    \"expires_at\": \"2026-10-23T10:01:00Z\",\n                    \"instructions\": {\n                      \"source_account_ref\": \"acct_example_01\",\n                      \"source_asset\": {\n                        \"network\": \"ethereum_sepolia\",\n                        \"currency\": \"USDC\",\n                        \"kind\": \"TOKEN_CONTRACT\",\n                        \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                      },\n                      \"destination\": {\n                        \"asset\": {\n                          \"network\": \"ethereum_sepolia\",\n                          \"currency\": \"USDC\",\n                          \"kind\": \"TOKEN_CONTRACT\",\n                          \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                        },\n                        \"address\": \"0x2222222222222222222222222222222222222222\"\n                      }\n                    }\n                  },\n                  \"executions\": [\n                    {\n                      \"execution_ref\": \"execution_example_01\",\n                      \"network\": \"ethereum_sepolia\",\n                      \"status\": \"CONFIRMING\",\n                      \"transaction_hash\": \"0x4444444444444444444444444444444444444444444444444444444444444444\"\n                    }\n                  ],\n                  \"bridges\": [],\n                  \"created_at\": \"2026-10-23T10:00:02Z\",\n                  \"updated_at\": \"2026-10-23T10:01:10Z\"\n                }\n              }\n            }\n          }\n        },\n        \"responses\": {\n          \"204\": {\n            \"description\": \"Verified event durably accepted or already processed.\"\n          },\n          \"400\": {\n            \"description\": \"Malformed event.\"\n          },\n          \"401\": {\n            \"description\": \"Signature or timestamp verification failed.\"\n          },\n          \"503\": {\n            \"description\": \"Receiver could not durably accept; sender must retry with the same event ID and a new delivery identity/signature.\"\n          }\n        },\n        \"x-implementation\": \"api/processing/internal/publicapi/notifications.go: transactional task creation from the retained resource/effect projection; the original callback sender provides persisted retry, Vault-only HMAC signing and restart recovery.\"\n      }\n    },\n    \"depositUpdated\": {\n      \"post\": {\n        \"operationId\": \"depositUpdated\",\n        \"tags\": [\n          \"Notifications\"\n        ],\n        \"summary\": \"Deposit updates\",\n        \"description\": \"Resource update only. Pending/swept deposits are not financial credits. At-least-once delivery may arrive out of order. Verify the HMAC and timestamp before accepting, durably record/deduplicate the event, then return 2xx. Subscription setup is out-of-band. Compare deposit_ref/version and immutable receiving snapshot. Recover missed observations or later revisions with GET /v1/deposits; a pending/swept deposit does not authorize credit. Delivery is at least once and may be out of order. Acknowledge only after durable inbox storage; automatic retries use backoff after timeout/non-2xx, then require administrative replay or read recovery when exhausted. Numeric timeout/backoff/retry horizon and read/inbox retention must be verified and published before production; no numeric delivery SLA is asserted here. Publish only after its resource and its authorized deposit/transfer origins can be read by the same principal. A financial event may precede a resource-update webhook; resolve its origin by GET rather than waiting for notification order.\",\n        \"security\": [],\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"InboundDeposit durable state update\",\n        \"parameters\": [\n          {\n            \"$ref\": \"#/components/parameters/CallbackTimestamp\"\n          },\n          {\n            \"$ref\": \"#/components/parameters/CallbackSignature\"\n          },\n          {\n            \"$ref\": \"#/components/parameters/CallbackAlgorithm\"\n          }\n        ],\n        \"requestBody\": {\n          \"required\": true,\n          \"content\": {\n            \"application/json\": {\n              \"schema\": {\n                \"$ref\": \"#/components/schemas/DepositEvent\"\n              },\n              \"example\": {\n                \"event_id\": \"event_depositUpdated_01\",\n                \"delivery_id\": \"delivery_attempt_01\",\n                \"occurred_at\": \"2026-10-23T10:00:05Z\",\n                \"type\": \"deposit.updated\",\n                \"data\": {\n                  \"deposit_ref\": \"deposit_example_01\",\n                  \"inbound_ref\": \"inbound_example_01\",\n                  \"account_ref\": \"acct_example_01\",\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  },\n                  \"amount\": \"125.50\",\n                  \"state\": \"AML_PENDING\",\n                  \"created_at\": \"2026-10-23T10:00:00Z\",\n                  \"updated_at\": \"2026-10-23T10:00:05Z\",\n                  \"receiving\": {\n                    \"receiving_address_ref\": \"route_example_01\",\n                    \"route_version\": \"1\",\n                    \"address\": \"0x3333333333333333333333333333333333333333\",\n                    \"captured_at\": \"2026-10-23T10:00:00Z\"\n                  },\n                  \"version\": \"2\"\n                }\n              }\n            }\n          }\n        },\n        \"responses\": {\n          \"204\": {\n            \"description\": \"Verified event durably accepted or already processed.\"\n          },\n          \"400\": {\n            \"description\": \"Malformed event.\"\n          },\n          \"401\": {\n            \"description\": \"Signature or timestamp verification failed.\"\n          },\n          \"503\": {\n            \"description\": \"Receiver could not durably accept; sender must retry with the same event ID and a new delivery identity/signature.\"\n          }\n        },\n        \"x-implementation\": \"api/processing/internal/publicapi/notifications.go: transactional task creation from the retained resource/effect projection; the original callback sender provides persisted retry, Vault-only HMAC signing and restart recovery.\"\n      }\n    },\n    \"financialEffectCommitted\": {\n      \"post\": {\n        \"operationId\": \"financialEffectCommitted\",\n        \"tags\": [\n          \"Notifications\"\n        ],\n        \"summary\": \"Financial effects\",\n        \"description\": \"Produced only from a committed Accounting receipt/outbox. Deduplicate data.effect_ref atomically with the consumer posting. An Accounting outage after blockchain confirmation must delay this notification; chain status is insufficient. At-least-once delivery may arrive out of order. Verify the HMAC and timestamp before accepting, durably record/deduplicate the event, then return 2xx. Subscription setup is out-of-band. Recover the identical canonical effect through GET /v1/financial-effects/{effect_ref}, or discover missed effects with GET /v1/financial-effects. Commit local effect_ref deduplication with the posting in your application; event/delivery IDs alone do not deduplicate a financial posting. Delivery is at least once and may be out of order. Acknowledge only after durable inbox storage; automatic retries use backoff after timeout/non-2xx, then require administrative replay or read recovery when exhausted. Numeric timeout/backoff/retry horizon and read/inbox retention must be verified and published before production; no numeric delivery SLA is asserted here. Publish only after its resource and its authorized deposit/transfer origins can be read by the same principal. A financial event may precede a resource-update webhook; resolve its origin by GET rather than waiting for notification order.\",\n        \"security\": [],\n        \"x-status\": \"implemented\",\n        \"x-release-gate\": \"milestone-4-production-operations-and-client-integration-acceptance\",\n        \"x-source\": \"Ledger executePostingPlanSnapshot committed result + transactional outbox\",\n        \"parameters\": [\n          {\n            \"$ref\": \"#/components/parameters/CallbackTimestamp\"\n          },\n          {\n            \"$ref\": \"#/components/parameters/CallbackSignature\"\n          },\n          {\n            \"$ref\": \"#/components/parameters/CallbackAlgorithm\"\n          }\n        ],\n        \"requestBody\": {\n          \"required\": true,\n          \"content\": {\n            \"application/json\": {\n              \"schema\": {\n                \"$ref\": \"#/components/schemas/FinancialEvent\"\n              },\n              \"example\": {\n                \"event_id\": \"event_financialEffectCommitted_01\",\n                \"delivery_id\": \"delivery_attempt_01\",\n                \"occurred_at\": \"2026-10-23T10:02:00Z\",\n                \"type\": \"financial_effect.committed\",\n                \"data\": {\n                  \"effect_ref\": \"journal_example_01:line_example_01\",\n                  \"journal_entry_ref\": \"journal_example_01\",\n                  \"posting_line_ref\": \"line_example_01\",\n                  \"account_ref\": \"acct_example_01\",\n                  \"money\": {\n                    \"amount\": \"-125.50\",\n                    \"currency\": \"USDC\"\n                  },\n                  \"committed_at\": \"2026-10-23T10:02:00Z\",\n                  \"origin\": {\n                    \"type\": \"TRANSFER\",\n                    \"transfer_ref\": \"transfer_example_01\",\n                    \"client_reference\": \"merchant-withdrawal-2026-001\"\n                  },\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  },\n                  \"line_type\": \"MAIN\"\n                }\n              }\n            }\n          }\n        },\n        \"responses\": {\n          \"204\": {\n            \"description\": \"Verified event durably accepted or already processed.\"\n          },\n          \"400\": {\n            \"description\": \"Malformed event.\"\n          },\n          \"401\": {\n            \"description\": \"Signature or timestamp verification failed.\"\n          },\n          \"503\": {\n            \"description\": \"Receiver could not durably accept; sender must retry with the same event ID and a new delivery identity/signature.\"\n          }\n        },\n        \"x-implementation\": \"api/processing/internal/publicapi/notifications.go: transactional task creation from the retained resource/effect projection; the original callback sender provides persisted retry, Vault-only HMAC signing and restart recovery.\"\n      }\n    }\n  },\n  \"components\": {\n    \"securitySchemes\": {\n      \"TenantBearer\": {\n        \"type\": \"http\",\n        \"scheme\": \"bearer\",\n        \"description\": \"Tenant-scoped service credential. The trusted boundary derives tenant/actor identity and enforces resource ownership. Token issuance is outside this draft; no login or arbitrary tenant-selection API is defined.\"\n      }\n    },\n    \"parameters\": {\n      \"IdempotencyKey\": {\n        \"name\": \"Idempotency-Key\",\n        \"in\": \"header\",\n        \"required\": true,\n        \"description\": \"Persist the key and exact request before sending. Scope is authenticated tenant + operation + key. Same key/payload, including simultaneous requests, must resolve to the same durable resource/result, or return 409 SUBMISSION_IN_PROGRESS with retry of this same request/key. Different payload returns 409 IDEMPOTENCY_CONFLICT. Accepted transfer identity is resolved before quote-expiry validation; replay returns its frozen snapshot. Never reuse an expired retained identity: return 410 IDEMPOTENCY_RECORD_EXPIRED and reconcile. The current implementation applies no automatic TTL to public commands/claims; explicit retirement preserves the execution fence. Production archival/retention agreements remain required; see /guide.md#installed-retention-and-delivery-policy. Request equality uses the stored validated JSON field values and presence, not HTTP bytes: object property order/JSON whitespace are ignored; decimal string lexemes and supplied-vs-omitted fields are preserved. Thus \\\"100\\\" and \\\"100.00\\\" are different request values (JSON numbers are invalid for amounts). Omitted catalog-resolved contract_address and an explicitly supplied address are different under the same key, even if they resolved to the same asset. Replay compares the original stored input/comparison policy before current catalog resolution or precision checks, while enforcing current authorization; accepted terms are never re-resolved/repriced. Duplicate JSON member names, unknown request fields and disallowed nulls are rejected. See /guide.md#idempotency-request-comparison and the implemented replay evidence.\",\n        \"schema\": {\n          \"type\": \"string\",\n          \"description\": \"Persist with the exact request before sending. Preserve on retry. Scope: authenticated tenant and operation. No API-side key regeneration.\",\n          \"minLength\": 1,\n          \"examples\": [\n            \"merchant-request-2026-001\"\n          ]\n        }\n      },\n      \"Cursor\": {\n        \"name\": \"cursor\",\n        \"in\": \"query\",\n        \"required\": false,\n        \"description\": \"Opaque cursor issued by this resource owner. Do not reuse tokens across resources or tenants.\",\n        \"schema\": {\n          \"type\": \"string\"\n        }\n      },\n      \"Limit\": {\n        \"name\": \"limit\",\n        \"in\": \"query\",\n        \"required\": false,\n        \"description\": \"Maximum page size; the boundary validates this transport limit before forwarding.\",\n        \"schema\": {\n          \"type\": \"integer\",\n          \"minimum\": 1,\n          \"maximum\": 100,\n          \"default\": 25\n        }\n      },\n      \"CallbackTimestamp\": {\n        \"name\": \"X-Callback-Timestamp\",\n        \"in\": \"header\",\n        \"required\": true,\n        \"description\": \"UTC RFC3339Nano delivery timestamp. Verify a five-minute replay window.\",\n        \"schema\": {\n          \"$ref\": \"#/components/schemas/Timestamp\"\n        }\n      },\n      \"CallbackSignature\": {\n        \"name\": \"X-Callback-Signature\",\n        \"in\": \"header\",\n        \"required\": true,\n        \"description\": \"Lowercase hex HMAC-SHA256(secret, UTF8(timestamp + \\\".\\\") || rawBody). Verify before parsing, using constant-time comparison.\",\n        \"schema\": {\n          \"type\": \"string\",\n          \"pattern\": \"^[0-9a-f]{64}$\"\n        }\n      },\n      \"CallbackAlgorithm\": {\n        \"name\": \"X-Callback-Algorithm\",\n        \"in\": \"header\",\n        \"required\": true,\n        \"description\": \"Existing callback HMAC signing profile. Independent webhook secrets are provisioned outside this API.\",\n        \"schema\": {\n          \"type\": \"string\",\n          \"const\": \"hmac-sha256\"\n        }\n      },\n      \"RecoveryCursor\": {\n        \"name\": \"cursor\",\n        \"in\": \"query\",\n        \"required\": false,\n        \"description\": \"Omit cursor to read retained history. Continue using next_cursor, including after has_more=false or an empty page. Persist a returned cursor only after durably handling every item. Tokens bind the authenticated tenant, this surface, and unchanged filters; malformed/mismatched tokens return 400, expired or unrecoverable checkpoints return 410 and require reconciliation. Never silently reset. Cursor order is durable projection commit order, not created_at/event_time, so late commits and state updates remain discoverable. Projection recovery must not advance past unprojected committed source changes. Current recovery cursors expire after 24 hours; retained revisions/effects have no automatic TTL. Production retention, archival and replay agreements remain release gates; see /guide.md#installed-retention-and-delivery-policy.\",\n        \"schema\": {\n          \"type\": \"string\",\n          \"minLength\": 1\n        }\n      }\n    },\n    \"responses\": {\n      \"BadRequest\": {\n        \"description\": \"Malformed JSON or invalid HTTP parameter.\",\n        \"content\": {\n          \"application/json\": {\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Error\"\n            },\n            \"example\": {\n              \"code\": \"INVALID_REQUEST\",\n              \"message\": \"Malformed JSON or invalid HTTP parameter.\",\n              \"request_id\": \"request_example_01\"\n            }\n          }\n        }\n      },\n      \"Unauthorized\": {\n        \"description\": \"Missing or invalid bearer credential.\",\n        \"content\": {\n          \"application/json\": {\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Error\"\n            },\n            \"example\": {\n              \"code\": \"UNAUTHENTICATED\",\n              \"message\": \"Missing or invalid bearer credential.\",\n              \"request_id\": \"request_example_01\"\n            }\n          }\n        }\n      },\n      \"Forbidden\": {\n        \"description\": \"Authenticated principal lacks this operation permission.\",\n        \"content\": {\n          \"application/json\": {\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Error\"\n            },\n            \"example\": {\n              \"code\": \"PERMISSION_DENIED\",\n              \"message\": \"The principal cannot perform this operation.\",\n              \"request_id\": \"request_example_01\"\n            }\n          }\n        }\n      },\n      \"NotFound\": {\n        \"description\": \"Resource absent or outside the authenticated visibility scope.\",\n        \"content\": {\n          \"application/json\": {\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Error\"\n            },\n            \"example\": {\n              \"code\": \"NOT_FOUND\",\n              \"message\": \"Resource not found in the authorized scope.\",\n              \"request_id\": \"request_example_01\"\n            }\n          }\n        }\n      },\n      \"Conflict\": {\n        \"description\": \"Conflicting input, duplicate incompatible quote claim, unavailable execution admission, or unexecutable quote. Retry guidance is code-specific; never change a key to bypass an uncertain accepted request.\",\n        \"content\": {\n          \"application/json\": {\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Error\"\n            },\n            \"examples\": {\n              \"IDEMPOTENCY_CONFLICT\": {\n                \"summary\": \"IDEMPOTENCY CONFLICT\",\n                \"value\": {\n                  \"code\": \"IDEMPOTENCY_CONFLICT\",\n                  \"message\": \"The submission key belongs to different input.\",\n                  \"request_id\": \"request_example_01\"\n                }\n              },\n              \"SUBMISSION_IN_PROGRESS\": {\n                \"summary\": \"SUBMISSION IN PROGRESS\",\n                \"value\": {\n                  \"code\": \"SUBMISSION_IN_PROGRESS\",\n                  \"message\": \"The original request is being durably accepted; retry the exact original request and key.\",\n                  \"request_id\": \"request_example_01\"\n                }\n              },\n              \"QUOTE_CLAIM_CONFLICT\": {\n                \"summary\": \"QUOTE CLAIM CONFLICT\",\n                \"value\": {\n                  \"code\": \"QUOTE_CLAIM_CONFLICT\",\n                  \"message\": \"The quote intent is already claimed by incompatible execution. Reconcile the original transfer.\",\n                  \"request_id\": \"request_example_01\"\n                }\n              },\n              \"QUOTE_EXPIRED\": {\n                \"summary\": \"QUOTE EXPIRED\",\n                \"value\": {\n                  \"code\": \"QUOTE_EXPIRED\",\n                  \"message\": \"This unaccepted quote has expired; inspect existing transfers by quote_ref before requesting replacement terms.\",\n                  \"request_id\": \"request_example_01\"\n                }\n              },\n              \"QUOTE_NOT_EXECUTABLE\": {\n                \"summary\": \"QUOTE NOT EXECUTABLE\",\n                \"value\": {\n                  \"code\": \"QUOTE_NOT_EXECUTABLE\",\n                  \"message\": \"A blocked or indicative quote cannot be submitted.\",\n                  \"request_id\": \"request_example_01\"\n                }\n              },\n              \"LIQUIDITY_UNAVAILABLE\": {\n                \"summary\": \"LIQUIDITY UNAVAILABLE\",\n                \"value\": {\n                  \"code\": \"LIQUIDITY_UNAVAILABLE\",\n                  \"message\": \"A pre-admission policy check could not admit this operation. An already accepted operation keeps its original terms and identity while Arctex resolves liquidity.\",\n                  \"request_id\": \"request_example_01\"\n                }\n              },\n              \"COMPLIANCE_REVIEW_REQUIRED\": {\n                \"summary\": \"COMPLIANCE REVIEW REQUIRED\",\n                \"value\": {\n                  \"code\": \"COMPLIANCE_REVIEW_REQUIRED\",\n                  \"message\": \"Operator review is required; resubmission does not bypass screening.\",\n                  \"request_id\": \"request_example_01\"\n                }\n              }\n            }\n          }\n        }\n      },\n      \"InvalidInput\": {\n        \"description\": \"Semantic input/pre-admission rejection; inspect code. Post-acceptance financial/workflow failures are resource attention, not an HTTP rollback.\",\n        \"content\": {\n          \"application/json\": {\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Error\"\n            },\n            \"examples\": {\n              \"INVALID_INPUT\": {\n                \"summary\": \"INVALID INPUT\",\n                \"value\": {\n                  \"code\": \"INVALID_INPUT\",\n                  \"message\": \"The source/destination or supported asset is invalid.\",\n                  \"request_id\": \"request_example_01\"\n                }\n              },\n              \"INVALID_PRECISION\": {\n                \"summary\": \"INVALID PRECISION\",\n                \"value\": {\n                  \"code\": \"INVALID_PRECISION\",\n                  \"message\": \"The requested amount exceeds the authoritative asset accounting scale; input was not rounded.\",\n                  \"request_id\": \"request_example_01\"\n                }\n              },\n              \"UNSUPPORTED_QUOTE_COMBINATION\": {\n                \"summary\": \"UNSUPPORTED QUOTE COMBINATION\",\n                \"value\": {\n                  \"code\": \"UNSUPPORTED_QUOTE_COMBINATION\",\n                  \"message\": \"This route does not support the selected amount mode and fee treatment. No fallback mode was used.\",\n                  \"request_id\": \"request_example_01\"\n                }\n              },\n              \"INSUFFICIENT_FUNDS\": {\n                \"summary\": \"INSUFFICIENT FUNDS\",\n                \"value\": {\n                  \"code\": \"INSUFFICIENT_FUNDS\",\n                  \"message\": \"Accounting rejected reservation before transfer admission. After a previous uncertain response, reconcile that submission before retrying a new payment.\",\n                  \"request_id\": \"request_example_01\"\n                }\n              }\n            }\n          }\n        }\n      },\n      \"RateLimited\": {\n        \"description\": \"Request rate limited; retry according to Retry-After.\",\n        \"content\": {\n          \"application/json\": {\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Error\"\n            },\n            \"example\": {\n              \"code\": \"RATE_LIMITED\",\n              \"message\": \"Retry according to Retry-After.\",\n              \"request_id\": \"request_example_01\"\n            }\n          }\n        },\n        \"headers\": {\n          \"Retry-After\": {\n            \"description\": \"Suggested delay in seconds when known. Not permission to generate a new submission identity.\",\n            \"schema\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            }\n          }\n        }\n      },\n      \"Unavailable\": {\n        \"description\": \"Required Arctex dependency unavailable. Outcome may be uncertain: replay the original mutation with the original identity.\",\n        \"content\": {\n          \"application/json\": {\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Error\"\n            },\n            \"example\": {\n              \"code\": \"DEPENDENCY_UNAVAILABLE\",\n              \"message\": \"Outcome may be uncertain. Replay the original request with the original key.\",\n              \"request_id\": \"request_example_01\"\n            }\n          }\n        },\n        \"headers\": {\n          \"Retry-After\": {\n            \"description\": \"Suggested delay in seconds when known. Not permission to generate a new submission identity.\",\n            \"schema\": {\n              \"type\": \"integer\",\n              \"minimum\": 0\n            }\n          }\n        }\n      },\n      \"RecoveryCursorExpired\": {\n        \"description\": \"RECOVERY_CURSOR_EXPIRED: the checkpoint is expired or unrecoverable. Stop automatic cursor reset and reconcile retained resource/receipt identities with Arctex operations. Numeric supported retention must be published before production. Follow /guide.md#reconciliation-after-expiry: preserve inbox/dedup state, reconcile retained history to source receipts and an Arctex-issued replacement watermark, then resume with the same filters. A cursorless partial history is not proof of completeness.\",\n        \"content\": {\n          \"application/json\": {\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Error\"\n            },\n            \"example\": {\n              \"code\": \"RECOVERY_CURSOR_EXPIRED\",\n              \"message\": \"The checkpoint cannot be resumed. Reconcile retained receipts before establishing a new checkpoint.\",\n              \"request_id\": \"request_example_01\"\n            }\n          }\n        }\n      },\n      \"IdempotencyExpired\": {\n        \"description\": \"IDEMPOTENCY_RECORD_EXPIRED: the original submission identity is retired and cannot be replayed automatically. Reconcile the original durable resource with Arctex operations. Never reuse this key to start another execution. Numeric retention and retired-key fencing must be verified before production. Follow /guide.md#reconciliation-after-expiry: retrieve the known transfer or list by quote_ref; an empty result is not proof of non-execution. Arctex must resolve/fence every old execution/financial effect before any replacement payment is authorized. Keep the retired submission key blocked.\",\n        \"content\": {\n          \"application/json\": {\n            \"schema\": {\n              \"$ref\": \"#/components/schemas/Error\"\n            },\n            \"example\": {\n              \"code\": \"IDEMPOTENCY_RECORD_EXPIRED\",\n              \"message\": \"Reconcile the original submission before continuing; do not submit a new execution under this identity.\",\n              \"request_id\": \"request_example_01\"\n            }\n          }\n        }\n      }\n    },\n    \"schemas\": {\n      \"Reference\": {\n        \"type\": \"string\",\n        \"description\": \"Opaque identifier. Preserve exactly; do not infer ownership or resource type from its spelling.\",\n        \"minLength\": 1,\n        \"examples\": [\n          \"ref_example_01\"\n        ]\n      },\n      \"Decimal\": {\n        \"type\": \"string\",\n        \"description\": \"Exact base-10 decimal string; never a JSON number. Supported scale is validated against the authoritative catalog.\",\n        \"pattern\": \"^-?(0|[1-9][0-9]*)(\\\\.[0-9]+)?$\",\n        \"examples\": [\n          \"125.50\"\n        ]\n      },\n      \"PositiveDecimal\": {\n        \"type\": \"string\",\n        \"description\": \"Strictly positive exact amount. Zero is invalid. Asset precision is checked by the owner.\",\n        \"pattern\": \"^(?=.*[1-9])(0|[1-9][0-9]*)(\\\\.[0-9]+)?$\",\n        \"examples\": [\n          \"125.50\"\n        ]\n      },\n      \"Timestamp\": {\n        \"type\": \"string\",\n        \"description\": \"UTC RFC3339 timestamp.\",\n        \"format\": \"date-time\",\n        \"examples\": [\n          \"2026-10-23T10:00:00Z\"\n        ]\n      },\n      \"Version\": {\n        \"type\": \"string\",\n        \"description\": \"Existing owner-supplied nonnegative integer version, serialized as a string to preserve uint64 precision. Not a new sequence. Consumers retain the decimal string exactly or use arbitrary-precision integer handling for comparison; never decode through floating point. Version scopes are owner-specific, never global commit cutoffs.\",\n        \"pattern\": \"^[0-9]+$\",\n        \"examples\": [\n          \"7\",\n          \"9007199254740993\",\n          \"18446744073709551615\"\n        ]\n      },\n      \"Money\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"amount\": {\n            \"$ref\": \"#/components/schemas/Decimal\"\n          },\n          \"currency\": {\n            \"type\": \"string\",\n            \"description\": \"Currency/token code from the authoritative catalog.\",\n            \"examples\": [\n              \"USDC\"\n            ]\n          }\n        },\n        \"required\": [\n          \"amount\",\n          \"currency\"\n        ]\n      },\n      \"Asset\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"network\": {\n            \"type\": \"string\",\n            \"description\": \"Canonical network name from the Arctex catalog / chain configuration. Examples do not advertise enablement.\",\n            \"examples\": [\n              \"ethereum_sepolia\"\n            ]\n          },\n          \"currency\": {\n            \"type\": \"string\",\n            \"description\": \"Canonical token/native symbol.\",\n            \"examples\": [\n              \"USDC\"\n            ]\n          },\n          \"kind\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"NATIVE\",\n              \"TOKEN_CONTRACT\"\n            ]\n          },\n          \"contract_address\": {\n            \"type\": \"string\",\n            \"description\": \"Canonical token contract when present; validated against the catalog. The server resolves it when omitted for a token. Absent for native assets. Requests cannot override the catalog mapping. For idempotency comparison, omission is retained separately from the original catalog-resolved value; adding the resolved address on replay changes the request and conflicts.\"\n          }\n        },\n        \"required\": [\n          \"network\",\n          \"currency\",\n          \"kind\"\n        ],\n        \"description\": \"Network-specific chain asset. Catalog definitions remain authoritative; this structure is not an independent registry.\"\n      },\n      \"AssetCatalogEntry\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          },\n          \"accounting_scale\": {\n            \"type\": \"integer\",\n            \"minimum\": 0\n          },\n          \"display_scale\": {\n            \"type\": \"integer\",\n            \"minimum\": 0\n          }\n        },\n        \"required\": [\n          \"asset\",\n          \"accounting_scale\",\n          \"display_scale\"\n        ]\n      },\n      \"AssetList\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"catalog_revision\": {\n            \"type\": \"string\",\n            \"description\": \"Revision/hash of the active authoritative snapshot.\"\n          },\n          \"items\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/AssetCatalogEntry\"\n            }\n          }\n        },\n        \"required\": [\n          \"catalog_revision\",\n          \"items\"\n        ]\n      },\n      \"Account\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"account_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"account_code\": {\n            \"type\": \"string\",\n            \"description\": \"Existing account code, visible only within the authorized tenant scope.\"\n          },\n          \"status\": {\n            \"type\": \"string\",\n            \"description\": \"Existing financial account status; passed through from Accounting.\"\n          },\n          \"asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          }\n        },\n        \"required\": [\n          \"account_ref\",\n          \"status\",\n          \"asset\"\n        ],\n        \"description\": \"A financial account in Arctex One, separate from a physical custody wallet. Each account has one authoritative network-specific asset binding; a shared currency symbol does not permit implicit cross-network spending. Account provisioning and arbitrary posting remain authorized administrative operations.\"\n      },\n      \"AccountList\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"items\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/Account\"\n            }\n          }\n        },\n        \"required\": [\n          \"items\"\n        ]\n      },\n      \"Balance\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"account_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"balance\": {\n            \"$ref\": \"#/components/schemas/Money\"\n          },\n          \"version\": {\n            \"$ref\": \"#/components/schemas/Version\",\n            \"description\": \"Existing materialized account version. Several committed posting lines can share its value. It is not comparable to transfer/deposit versions, event time, journal version, or opaque recovery cursors.\"\n          },\n          \"last_updated\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"as_of\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          }\n        },\n        \"required\": [\n          \"account_ref\",\n          \"balance\",\n          \"asset\"\n        ],\n        \"description\": \"Current or historical Accounting balance, not a chain balance or a spend authorization. Historical responses omit version/last_updated and return as_of. The response reports committed balance; it does not expose available or reserved balance buckets. asset is the authoritative account liability identity and must match balance.currency; it cannot be inferred from the custody wallet currently backing it. Admission can precede reservation. Concurrent withdrawals are serialized/validated by authoritative Accounting hold/posting operations, never by a caller read-then-subtract calculation. Reservations and pending debits are not inferred from this balance. See /guide.md for release and failure rules. version is the materialized account balance version, not a journal/posting count or an effect-stream cursor. as_of is an event-time query, not a commit snapshot: later backdated postings can change a historical result. Reading this balance and draining financial effects independently does not establish a consistent accounting cutoff. Use the source-backed procedure in /guide.md#consistent-reconciliation-cutoff; no shared snapshot API is promised.\"\n      },\n      \"StatementSummary\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"entry_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"entry_type\": {\n            \"type\": \"string\",\n            \"description\": \"Existing journal entry type.\"\n          },\n          \"account_amount\": {\n            \"$ref\": \"#/components/schemas/Money\"\n          },\n          \"description\": {\n            \"type\": \"string\"\n          },\n          \"event_time\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"version\": {\n            \"$ref\": \"#/components/schemas/Version\"\n          },\n          \"visibility\": {\n            \"type\": \"string\",\n            \"description\": \"Existing account-scoped visibility classification.\"\n          },\n          \"asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          },\n          \"effect_refs\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            },\n            \"minItems\": 1,\n            \"uniqueItems\": true,\n            \"description\": \"Stable effects for this authorized account in the entry; retrieve each through /v1/financial-effects/{effect_ref}. No other accounts or balancing lines are exposed.\"\n          }\n        },\n        \"required\": [\n          \"entry_ref\",\n          \"entry_type\",\n          \"account_amount\",\n          \"event_time\",\n          \"version\",\n          \"asset\",\n          \"effect_refs\"\n        ]\n      },\n      \"PostingLine\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"account_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"money\": {\n            \"$ref\": \"#/components/schemas/Money\"\n          },\n          \"line_type\": {\n            \"$ref\": \"#/components/schemas/PostingClassification\"\n          },\n          \"posting_line_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"effect_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"origin\": {\n            \"$ref\": \"#/components/schemas/EffectOrigin\"\n          },\n          \"asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          }\n        },\n        \"required\": [\n          \"account_ref\",\n          \"money\",\n          \"line_type\",\n          \"posting_line_ref\",\n          \"effect_ref\",\n          \"origin\",\n          \"asset\"\n        ],\n        \"description\": \"Only lines visible within the requested account and authorized tenant scope. Stable identities and origin must match FinancialEffect, including distinct fee lines; current statement DTOs omit line identity, so enrich from committed receipts without changing journals. asset and money.currency match the authorized account.\"\n      },\n      \"StatementEntry\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"entry_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"entry_type\": {\n            \"type\": \"string\",\n            \"description\": \"Existing journal entry type.\"\n          },\n          \"account_amount\": {\n            \"$ref\": \"#/components/schemas/Money\"\n          },\n          \"description\": {\n            \"type\": \"string\"\n          },\n          \"event_time\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"version\": {\n            \"$ref\": \"#/components/schemas/Version\"\n          },\n          \"visibility\": {\n            \"type\": \"string\",\n            \"description\": \"Existing account-scoped visibility classification.\"\n          },\n          \"lines\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/PostingLine\"\n            }\n          },\n          \"asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          },\n          \"effect_refs\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/Reference\"\n            },\n            \"minItems\": 1,\n            \"uniqueItems\": true,\n            \"description\": \"Stable effects for this authorized account in the entry; retrieve each through /v1/financial-effects/{effect_ref}. No other accounts or balancing lines are exposed.\"\n          }\n        },\n        \"required\": [\n          \"entry_ref\",\n          \"entry_type\",\n          \"account_amount\",\n          \"event_time\",\n          \"version\",\n          \"lines\",\n          \"asset\",\n          \"effect_refs\"\n        ],\n        \"description\": \"Committed account-scoped journal projection. Compensations are separate entries; historical effects are immutable.\"\n      },\n      \"StatementList\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"items\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/StatementSummary\"\n            }\n          }\n        },\n        \"required\": [\n          \"items\"\n        ],\n        \"description\": \"Uses limit/offset, matching the existing statement API. There is no fabricated continuation cursor or total count.\"\n      },\n      \"Destination\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          },\n          \"address\": {\n            \"type\": \"string\",\n            \"description\": \"Chain-validated recipient address. A syntactically valid address does not prove recipient ownership.\",\n            \"minLength\": 1\n          },\n          \"memo\": {\n            \"type\": \"string\",\n            \"description\": \"Optional network-specific memo/tag; validated for the selected network.\"\n          }\n        },\n        \"required\": [\n          \"asset\",\n          \"address\"\n        ]\n      },\n      \"CreateReceivingAddress\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"account_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          }\n        },\n        \"required\": [\n          \"account_ref\",\n          \"asset\"\n        ],\n        \"description\": \"The boundary validates account ownership and asset compatibility, then reuses existing walletprovisioning and receiving-route coordination. It selects custody infrastructure internally.\"\n      },\n      \"ReceivingAddress\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"receiving_address_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"account_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          },\n          \"address\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          },\n          \"memo\": {\n            \"type\": \"string\"\n          },\n          \"status\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"ACTIVE\",\n              \"DISABLED\"\n            ]\n          },\n          \"version\": {\n            \"$ref\": \"#/components/schemas/Version\"\n          },\n          \"created_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          }\n        },\n        \"required\": [\n          \"receiving_address_ref\",\n          \"account_ref\",\n          \"asset\",\n          \"address\",\n          \"status\",\n          \"version\"\n        ],\n        \"description\": \"Tenant-visible receiving-route projection. The reference maps to the committed route. Multiple addresses may credit one account; execution pools remain shared and private. DISABLED stops future allocation/selection; it does not stop observation or erase attribution. Rotation creates a new receiving_address_ref. Late deposits at retained addresses still undergo the original owner recognition, screening, and exactly-once credit path; changing a current route cannot reassign historical deposits. Lifecycle changes use existing administrative controls, not new public mutation endpoints.\"\n      },\n      \"ReceivingAddressList\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"items\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/ReceivingAddress\"\n            }\n          },\n          \"next_cursor\": {\n            \"type\": \"string\",\n            \"description\": \"Opaque next page token; absent when the existing owner returns no next page.\"\n          }\n        },\n        \"required\": [\n          \"items\"\n        ]\n      },\n      \"Deposit\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"deposit_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"inbound_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"account_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          },\n          \"amount\": {\n            \"$ref\": \"#/components/schemas/PositiveDecimal\"\n          },\n          \"state\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"OBSERVED\",\n              \"REVIEW_REQUIRED\",\n              \"REJECTED\",\n              \"APPROVED_WAITING_SWEEP\",\n              \"SWEEP_IN_PROGRESS\",\n              \"SWEEP_RETRYABLE_FAILURE\",\n              \"CREDIT_READY\",\n              \"CREDIT_RETRYABLE_FAILURE\",\n              \"COMPLETED\",\n              \"AML_PENDING\",\n              \"REMEDIATION_IN_PROGRESS\"\n            ]\n          },\n          \"credit_journal_entry_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"created_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"updated_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"receiving\": {\n            \"$ref\": \"#/components/schemas/ReceivingAttribution\"\n          },\n          \"version\": {\n            \"$ref\": \"#/components/schemas/Version\"\n          },\n          \"attention\": {\n            \"$ref\": \"#/components/schemas/Attention\"\n          },\n          \"chain_evidence\": {\n            \"$ref\": \"#/components/schemas/ChainEvidence\"\n          }\n        },\n        \"required\": [\n          \"deposit_ref\",\n          \"inbound_ref\",\n          \"asset\",\n          \"amount\",\n          \"state\",\n          \"created_at\",\n          \"updated_at\",\n          \"receiving\",\n          \"version\"\n        ],\n        \"description\": \"An incoming digital asset deposit with immutable intake and revisioned receiving-address attribution. account_ref may be absent before financial attribution; once resolved, historical revisions retain their original values. Corrections append compensating effects rather than rewriting committed credit history. credit_journal_entry_ref appears only from a committed credit receipt. Observed, approved, swept, or chain-finalized does not mean credited. Retain the same deposit identity through re-observation and recovery. Replacement/reorg evidence is resolved by existing chain-specific recognition; do not derive a fresh public identity merely from a changed transaction hash. Producer deduplication uses chain-specific transfer/log identity, never a transaction hash alone. Missing legacy attribution requires an explicit adoption procedure; it cannot be fabricated from current routing. attention is mandatory for review, rejection, retry-failure, and remediation states. WAIT describes automatic recovery without implying customer credit or immediate execution. Receiving evidence may lack a route reference while attribution is unresolved. COMPLETED or the presence of credit_journal_entry_ref requires a resolved account_ref and receiving.receiving_address_ref backed by the committed credit. A route reference alone does not prove financial credit.\",\n        \"allOf\": [\n          {\n            \"if\": {\n              \"properties\": {\n                \"state\": {\n                  \"const\": \"COMPLETED\"\n                }\n              },\n              \"required\": [\n                \"state\"\n              ],\n              \"type\": \"object\",\n              \"additionalProperties\": true\n            },\n            \"then\": {\n              \"required\": [\n                \"credit_journal_entry_ref\",\n                \"account_ref\"\n              ],\n              \"properties\": {\n                \"receiving\": {\n                  \"required\": [\n                    \"receiving_address_ref\"\n                  ],\n                  \"type\": \"object\",\n                  \"additionalProperties\": true,\n                  \"properties\": {\n                    \"receiving_address_ref\": {\n                      \"$ref\": \"#/components/schemas/Reference\"\n                    }\n                  }\n                },\n                \"credit_journal_entry_ref\": {\n                  \"$ref\": \"#/components/schemas/Deposit/properties/credit_journal_entry_ref\"\n                },\n                \"account_ref\": {\n                  \"$ref\": \"#/components/schemas/Deposit/properties/account_ref\"\n                }\n              },\n              \"type\": \"object\",\n              \"additionalProperties\": true\n            }\n          },\n          {\n            \"if\": {\n              \"required\": [\n                \"credit_journal_entry_ref\"\n              ],\n              \"type\": \"object\",\n              \"additionalProperties\": true,\n              \"properties\": {\n                \"credit_journal_entry_ref\": {\n                  \"$ref\": \"#/components/schemas/Deposit/properties/credit_journal_entry_ref\"\n                }\n              }\n            },\n            \"then\": {\n              \"required\": [\n                \"account_ref\"\n              ],\n              \"properties\": {\n                \"receiving\": {\n                  \"required\": [\n                    \"receiving_address_ref\"\n                  ],\n                  \"type\": \"object\",\n                  \"additionalProperties\": true,\n                  \"properties\": {\n                    \"receiving_address_ref\": {\n                      \"$ref\": \"#/components/schemas/Reference\"\n                    }\n                  }\n                },\n                \"account_ref\": {\n                  \"$ref\": \"#/components/schemas/Deposit/properties/account_ref\"\n                }\n              },\n              \"type\": \"object\",\n              \"additionalProperties\": true\n            }\n          },\n          {\n            \"if\": {\n              \"type\": \"object\",\n              \"properties\": {\n                \"state\": {\n                  \"enum\": [\n                    \"REVIEW_REQUIRED\",\n                    \"REJECTED\",\n                    \"SWEEP_RETRYABLE_FAILURE\",\n                    \"CREDIT_RETRYABLE_FAILURE\",\n                    \"REMEDIATION_IN_PROGRESS\"\n                  ]\n                }\n              },\n              \"required\": [\n                \"state\"\n              ],\n              \"additionalProperties\": true\n            },\n            \"then\": {\n              \"type\": \"object\",\n              \"required\": [\n                \"attention\"\n              ],\n              \"properties\": {\n                \"attention\": {\n                  \"$ref\": \"#/components/schemas/Attention\"\n                }\n              },\n              \"additionalProperties\": true\n            }\n          }\n        ],\n        \"examples\": [\n          {\n            \"deposit_ref\": \"deposit_example_01\",\n            \"inbound_ref\": \"inbound_example_01\",\n            \"asset\": {\n              \"network\": \"ethereum_sepolia\",\n              \"currency\": \"USDC\",\n              \"kind\": \"TOKEN_CONTRACT\",\n              \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n            },\n            \"amount\": \"125.50\",\n            \"state\": \"REVIEW_REQUIRED\",\n            \"created_at\": \"2026-10-23T10:00:00Z\",\n            \"updated_at\": \"2026-10-23T10:00:05Z\",\n            \"receiving\": {\n              \"address\": \"0x3333333333333333333333333333333333333333\",\n              \"captured_at\": \"2026-10-23T10:00:00Z\"\n            },\n            \"version\": \"1\",\n            \"attention\": {\n              \"reason_code\": \"ROUTE_UNRESOLVED\",\n              \"message\": \"Receiving attribution requires operator review. No financial credit has been committed.\",\n              \"action\": \"CONTACT_SUPPORT\"\n            }\n          },\n          {\n            \"deposit_ref\": \"deposit_example_01\",\n            \"inbound_ref\": \"inbound_example_01\",\n            \"account_ref\": \"acct_example_01\",\n            \"asset\": {\n              \"network\": \"ethereum_sepolia\",\n              \"currency\": \"USDC\",\n              \"kind\": \"TOKEN_CONTRACT\",\n              \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n            },\n            \"amount\": \"125.50\",\n            \"state\": \"COMPLETED\",\n            \"created_at\": \"2026-10-23T10:00:00Z\",\n            \"updated_at\": \"2026-10-24T11:00:00Z\",\n            \"receiving\": {\n              \"receiving_address_ref\": \"route_example_01\",\n              \"route_version\": \"1\",\n              \"address\": \"0x3333333333333333333333333333333333333333\",\n              \"captured_at\": \"2026-10-23T10:00:00Z\"\n            },\n            \"version\": \"7\",\n            \"credit_journal_entry_ref\": \"journal_deposit_example_01\"\n          },\n          {\n            \"deposit_ref\": \"deposit_example_01\",\n            \"inbound_ref\": \"inbound_example_01\",\n            \"account_ref\": \"acct_example_01\",\n            \"asset\": {\n              \"network\": \"ethereum_sepolia\",\n              \"currency\": \"USDC\",\n              \"kind\": \"TOKEN_CONTRACT\",\n              \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n            },\n            \"amount\": \"125.50\",\n            \"state\": \"COMPLETED\",\n            \"created_at\": \"2026-10-23T10:00:00Z\",\n            \"updated_at\": \"2026-10-24T11:00:00Z\",\n            \"receiving\": {\n              \"receiving_address_ref\": \"route_example_01\",\n              \"route_version\": \"1\",\n              \"address\": \"0x3333333333333333333333333333333333333333\",\n              \"captured_at\": \"2026-10-23T10:00:00Z\"\n            },\n            \"version\": \"7\",\n            \"credit_journal_entry_ref\": \"journal_deposit_example_01\",\n            \"chain_evidence\": {\n              \"network\": \"ethereum_sepolia\",\n              \"transaction_hash\": \"0x5555555555555555555555555555555555555555555555555555555555555555\"\n            }\n          }\n        ]\n      },\n      \"BridgeSelection\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"route_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"mode\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"STANDARD\",\n              \"FAST\"\n            ]\n          }\n        },\n        \"required\": [\n          \"route_ref\",\n          \"mode\"\n        ],\n        \"description\": \"Existing public bridge selection. Server resolves wallets, destination mint controls, gas, signing, and finality. Invalid/ineligible route choices are rejected.\"\n      },\n      \"QuoteRequest\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"source_account_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"source_asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          },\n          \"destination\": {\n            \"$ref\": \"#/components/schemas/Destination\"\n          },\n          \"amount\": {\n            \"$ref\": \"#/components/schemas/PositiveDecimal\",\n            \"description\": \"Exact user amount in source_asset units for FIX_SOURCE, destination.asset units for FIX_RECEIVED. FIX_SOURCE+ADD_TO_SOURCE fixes source principal before charges; FIX_SOURCE+DEDUCT_FROM_DESTINATION fixes total source budget; FIX_RECEIVED fixes recipient net amount for either treatment, with the required gross-up when supported by the selected route. Route-specific support is mandatory; unsupported/unverified combinations return UNSUPPORTED_QUOTE_COMBINATION and must not fall back to another mode. See /guide.md for supported quote combinations. The installed ordinary route rejects FIX_RECEIVED with DEDUCT_FROM_DESTINATION; the existing USDC bridge supports that combination through its gross-up calculator.\"\n          },\n          \"settlement_mode\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"FIX_SOURCE\",\n              \"FIX_RECEIVED\"\n            ]\n          },\n          \"fee_treatment\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"ADD_TO_SOURCE\",\n              \"DEDUCT_FROM_DESTINATION\"\n            ]\n          },\n          \"bridge\": {\n            \"$ref\": \"#/components/schemas/BridgeSelection\"\n          },\n          \"client_reference\": {\n            \"type\": \"string\",\n            \"description\": \"Optional caller correlation; not an execution or idempotency identity.\"\n          }\n        },\n        \"required\": [\n          \"source_account_ref\",\n          \"source_asset\",\n          \"destination\",\n          \"amount\",\n          \"settlement_mode\",\n          \"fee_treatment\"\n        ],\n        \"description\": \"A quote request for a supported digital asset transfer. Explicit source and destination networks prevent implicit asset conversion. No fiat, FX product, physical wallet selection, or caller-supplied fee amounts. Excess input precision is rejected, not rounded. Only supported routes/mode combinations may produce EXECUTABLE economics. Each supported route must satisfy the requested amount meaning; unsupported combinations are rejected. Replaying a key compares original validated input and field presence, including exact decimal strings; do not fill a previously omitted token contract or reformat amount on retry. Catalog resolution for an accepted key is frozen.\"\n      },\n      \"Fee\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"code\": {\n            \"type\": \"string\",\n            \"description\": \"Existing derived fee code.\"\n          },\n          \"amount\": {\n            \"$ref\": \"#/components/schemas/Money\"\n          },\n          \"description\": {\n            \"type\": \"string\"\n          },\n          \"asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          }\n        },\n        \"required\": [\n          \"code\",\n          \"amount\",\n          \"asset\"\n        ],\n        \"description\": \"Arctex processing fees and protocol costs only. your customer pricing is excluded. Fee payer/treatment follows the accepted quote, not an API-side recalculation. asset identifies the charge liability network and must agree with amount.currency. Include only tenant-chargeable accepted fees; platform native gas is not automatically billable. Distinct currencies/assets cannot be added directly. Refund/charge timing follows existing committed posting and compensation rules; no universal refund promise is inferred from transfer state. Route fee-policy verification is required before publication.\"\n      },\n      \"Quote\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"quote_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"state\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"INDICATIVE\",\n              \"EXECUTABLE\",\n              \"BLOCKED\",\n              \"EXPIRED\"\n            ]\n          },\n          \"block_reason\": {\n            \"type\": \"string\",\n            \"description\": \"Existing QuoteBlockReason, when applicable.\"\n          },\n          \"principal\": {\n            \"$ref\": \"#/components/schemas/Money\",\n            \"description\": \"Existing authoritative TransferPrincipalAmount/DebitAmount projection in source units. It is not automatically the payer total, recipient net, or bridge gross burn. Do not recompute it from other fields; verify per-route mapping.\"\n          },\n          \"recipient_amount\": {\n            \"$ref\": \"#/components/schemas/Money\",\n            \"description\": \"Quoted recipient net in destination units, after applicable destination deductions. For FIX_RECEIVED this equals the requested amount; otherwise use the authoritative quoted result.\"\n          },\n          \"payer_total_debit\": {\n            \"$ref\": \"#/components/schemas/Money\",\n            \"description\": \"Accepted total customer debit in source units, including chargeable fees represented in those units. An unpriced native gas cost must not become an uncapped extra tenant debit. No implicit currency/network conversion.\"\n          },\n          \"fees\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/Fee\"\n            }\n          },\n          \"fee_treatment\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"ADD_TO_SOURCE\",\n              \"DEDUCT_FROM_DESTINATION\"\n            ]\n          },\n          \"settlement_mode\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"FIX_SOURCE\",\n              \"FIX_RECEIVED\"\n            ]\n          },\n          \"readiness\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"LIQUIDITY_READY\",\n              \"LIQUIDITY_OBTAINABLE\",\n              \"INDICATIVE\"\n            ]\n          },\n          \"priced_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"expires_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"instructions\": {\n            \"$ref\": \"#/components/schemas/TransferInstructions\"\n          }\n        },\n        \"required\": [\n          \"quote_ref\",\n          \"state\",\n          \"fee_treatment\",\n          \"settlement_mode\",\n          \"readiness\",\n          \"instructions\"\n        ],\n        \"description\": \"Quote creation does not reserve funds or admit execution. EXECUTABLE requires strictly positive principal, recipient and payer amounts, complete authoritative economics, priced_at/expires_at, and non-indicative readiness. BLOCKED requires block_reason. INDICATIVE/BLOCKED may omit the entire economics group when reliable pricing is unavailable; absence means unknown, never zero/free. If any group field is present, all group fields must be present. Indicative estimates are not acceptable for submission. An EXPIRED result can retain its original complete snapshot but cannot admit a new transfer. Accepted Transfer.quote always retains its EXECUTABLE snapshot and original expiry; wall-clock expiry does not change accepted terms. See /guide.md for amount-mode matrix, precision, fees, and supported-route evidence gates.\",\n        \"dependentRequired\": {\n          \"principal\": [\n            \"recipient_amount\",\n            \"payer_total_debit\",\n            \"fees\",\n            \"priced_at\",\n            \"expires_at\"\n          ],\n          \"recipient_amount\": [\n            \"principal\",\n            \"payer_total_debit\",\n            \"fees\",\n            \"priced_at\",\n            \"expires_at\"\n          ],\n          \"payer_total_debit\": [\n            \"principal\",\n            \"recipient_amount\",\n            \"fees\",\n            \"priced_at\",\n            \"expires_at\"\n          ],\n          \"fees\": [\n            \"principal\",\n            \"recipient_amount\",\n            \"payer_total_debit\",\n            \"priced_at\",\n            \"expires_at\"\n          ],\n          \"priced_at\": [\n            \"principal\",\n            \"recipient_amount\",\n            \"payer_total_debit\",\n            \"fees\",\n            \"expires_at\"\n          ],\n          \"expires_at\": [\n            \"principal\",\n            \"recipient_amount\",\n            \"payer_total_debit\",\n            \"fees\",\n            \"priced_at\"\n          ]\n        },\n        \"allOf\": [\n          {\n            \"if\": {\n              \"type\": \"object\",\n              \"properties\": {\n                \"state\": {\n                  \"const\": \"EXECUTABLE\"\n                }\n              },\n              \"required\": [\n                \"state\"\n              ],\n              \"additionalProperties\": true\n            },\n            \"then\": {\n              \"type\": \"object\",\n              \"required\": [\n                \"principal\",\n                \"recipient_amount\",\n                \"payer_total_debit\",\n                \"fees\",\n                \"priced_at\",\n                \"expires_at\"\n              ],\n              \"properties\": {\n                \"principal\": {\n                  \"allOf\": [\n                    {\n                      \"$ref\": \"#/components/schemas/Money\",\n                      \"description\": \"Existing authoritative TransferPrincipalAmount/DebitAmount projection in source units. It is not automatically the payer total, recipient net, or bridge gross burn. Do not recompute it from other fields; verify per-route mapping.\"\n                    },\n                    {\n                      \"type\": \"object\",\n                      \"properties\": {\n                        \"amount\": {\n                          \"$ref\": \"#/components/schemas/PositiveDecimal\"\n                        }\n                      },\n                      \"required\": [\n                        \"amount\"\n                      ],\n                      \"additionalProperties\": true\n                    }\n                  ]\n                },\n                \"recipient_amount\": {\n                  \"allOf\": [\n                    {\n                      \"$ref\": \"#/components/schemas/Money\",\n                      \"description\": \"Quoted recipient net in destination units, after applicable destination deductions. For FIX_RECEIVED this equals the requested amount; otherwise use the authoritative quoted result.\"\n                    },\n                    {\n                      \"type\": \"object\",\n                      \"properties\": {\n                        \"amount\": {\n                          \"$ref\": \"#/components/schemas/PositiveDecimal\"\n                        }\n                      },\n                      \"required\": [\n                        \"amount\"\n                      ],\n                      \"additionalProperties\": true\n                    }\n                  ]\n                },\n                \"payer_total_debit\": {\n                  \"allOf\": [\n                    {\n                      \"$ref\": \"#/components/schemas/Money\",\n                      \"description\": \"Accepted total customer debit in source units, including chargeable fees represented in those units. An unpriced native gas cost must not become an uncapped extra tenant debit. No implicit currency/network conversion.\"\n                    },\n                    {\n                      \"type\": \"object\",\n                      \"properties\": {\n                        \"amount\": {\n                          \"$ref\": \"#/components/schemas/PositiveDecimal\"\n                        }\n                      },\n                      \"required\": [\n                        \"amount\"\n                      ],\n                      \"additionalProperties\": true\n                    }\n                  ]\n                },\n                \"fees\": {\n                  \"type\": \"array\",\n                  \"items\": {\n                    \"$ref\": \"#/components/schemas/Fee\"\n                  }\n                },\n                \"priced_at\": {\n                  \"$ref\": \"#/components/schemas/Timestamp\"\n                },\n                \"expires_at\": {\n                  \"$ref\": \"#/components/schemas/Timestamp\"\n                },\n                \"readiness\": {\n                  \"type\": \"string\",\n                  \"enum\": [\n                    \"LIQUIDITY_READY\",\n                    \"LIQUIDITY_OBTAINABLE\"\n                  ]\n                }\n              },\n              \"additionalProperties\": true\n            }\n          },\n          {\n            \"if\": {\n              \"type\": \"object\",\n              \"properties\": {\n                \"state\": {\n                  \"const\": \"BLOCKED\"\n                }\n              },\n              \"required\": [\n                \"state\"\n              ],\n              \"additionalProperties\": true\n            },\n            \"then\": {\n              \"type\": \"object\",\n              \"required\": [\n                \"block_reason\"\n              ],\n              \"properties\": {\n                \"block_reason\": {\n                  \"type\": \"string\",\n                  \"minLength\": 1\n                }\n              },\n              \"additionalProperties\": true\n            }\n          }\n        ],\n        \"examples\": [\n          {\n            \"quote_ref\": \"quote_example_01\",\n            \"state\": \"EXECUTABLE\",\n            \"principal\": {\n              \"amount\": \"125.50\",\n              \"currency\": \"USDC\"\n            },\n            \"recipient_amount\": {\n              \"amount\": \"125.50\",\n              \"currency\": \"USDC\"\n            },\n            \"payer_total_debit\": {\n              \"amount\": \"126.00\",\n              \"currency\": \"USDC\"\n            },\n            \"fees\": [\n              {\n                \"code\": \"mc_processing\",\n                \"amount\": {\n                  \"amount\": \"0.50\",\n                  \"currency\": \"USDC\"\n                },\n                \"description\": \"Illustrative Arctex processing fee\",\n                \"asset\": {\n                  \"network\": \"ethereum_sepolia\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                }\n              }\n            ],\n            \"fee_treatment\": \"ADD_TO_SOURCE\",\n            \"settlement_mode\": \"FIX_SOURCE\",\n            \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n            \"priced_at\": \"2026-10-23T10:00:00Z\",\n            \"expires_at\": \"2026-10-23T10:01:00Z\",\n            \"instructions\": {\n              \"source_account_ref\": \"acct_example_01\",\n              \"source_asset\": {\n                \"network\": \"ethereum_sepolia\",\n                \"currency\": \"USDC\",\n                \"kind\": \"TOKEN_CONTRACT\",\n                \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n              },\n              \"destination\": {\n                \"asset\": {\n                  \"network\": \"ethereum_sepolia\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"address\": \"0x2222222222222222222222222222222222222222\"\n              }\n            }\n          },\n          {\n            \"quote_ref\": \"quote_example_01\",\n            \"state\": \"BLOCKED\",\n            \"fee_treatment\": \"ADD_TO_SOURCE\",\n            \"settlement_mode\": \"FIX_SOURCE\",\n            \"readiness\": \"INDICATIVE\",\n            \"instructions\": {\n              \"source_account_ref\": \"acct_example_01\",\n              \"source_asset\": {\n                \"network\": \"ethereum_sepolia\",\n                \"currency\": \"USDC\",\n                \"kind\": \"TOKEN_CONTRACT\",\n                \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n              },\n              \"destination\": {\n                \"asset\": {\n                  \"network\": \"ethereum_sepolia\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"address\": \"0x2222222222222222222222222222222222222222\"\n              }\n            },\n            \"block_reason\": \"PRICING_UNAVAILABLE\"\n          },\n          {\n            \"quote_ref\": \"quote_example_01\",\n            \"state\": \"INDICATIVE\",\n            \"fee_treatment\": \"ADD_TO_SOURCE\",\n            \"settlement_mode\": \"FIX_SOURCE\",\n            \"readiness\": \"INDICATIVE\",\n            \"instructions\": {\n              \"source_account_ref\": \"acct_example_01\",\n              \"source_asset\": {\n                \"network\": \"ethereum_sepolia\",\n                \"currency\": \"USDC\",\n                \"kind\": \"TOKEN_CONTRACT\",\n                \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n              },\n              \"destination\": {\n                \"asset\": {\n                  \"network\": \"ethereum_sepolia\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"address\": \"0x2222222222222222222222222222222222222222\"\n              }\n            }\n          },\n          {\n            \"quote_ref\": \"bridge_fix_source_add_to_source\",\n            \"state\": \"EXECUTABLE\",\n            \"principal\": {\n              \"amount\": \"100\",\n              \"currency\": \"USDC\"\n            },\n            \"recipient_amount\": {\n              \"amount\": \"100\",\n              \"currency\": \"USDC\"\n            },\n            \"payer_total_debit\": {\n              \"amount\": \"102.010102\",\n              \"currency\": \"USDC\"\n            },\n            \"fees\": [\n              {\n                \"code\": \"mc_processing\",\n                \"amount\": {\n                  \"amount\": \"1\",\n                  \"currency\": \"USDC\"\n                },\n                \"asset\": {\n                  \"network\": \"ethereum_mainnet\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"description\": \"Illustrative bridge test commercial fee.\"\n              },\n              {\n                \"code\": \"bridge_protocol\",\n                \"amount\": {\n                  \"amount\": \"1.010102\",\n                  \"currency\": \"USDC\"\n                },\n                \"asset\": {\n                  \"network\": \"arbitrum_one\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"description\": \"Illustrative bridge protocol cost at 100 basis points, explicitly accounted across the accepted bridge route.\"\n              }\n            ],\n            \"fee_treatment\": \"ADD_TO_SOURCE\",\n            \"settlement_mode\": \"FIX_SOURCE\",\n            \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n            \"priced_at\": \"2026-10-23T10:00:00Z\",\n            \"expires_at\": \"2026-10-23T10:01:00Z\",\n            \"instructions\": {\n              \"source_account_ref\": \"acct_example_01\",\n              \"source_asset\": {\n                \"network\": \"ethereum_mainnet\",\n                \"currency\": \"USDC\",\n                \"kind\": \"TOKEN_CONTRACT\",\n                \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n              },\n              \"destination\": {\n                \"asset\": {\n                  \"network\": \"arbitrum_one\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"address\": \"0x2222222222222222222222222222222222222222\"\n              },\n              \"bridge\": {\n                \"route_ref\": \"illustrative_bridge_route_01\",\n                \"mode\": \"FAST\"\n              }\n            }\n          },\n          {\n            \"quote_ref\": \"bridge_fix_source_deduct_from_destination\",\n            \"state\": \"EXECUTABLE\",\n            \"principal\": {\n              \"amount\": \"100\",\n              \"currency\": \"USDC\"\n            },\n            \"recipient_amount\": {\n              \"amount\": \"98.01\",\n              \"currency\": \"USDC\"\n            },\n            \"payer_total_debit\": {\n              \"amount\": \"100\",\n              \"currency\": \"USDC\"\n            },\n            \"fees\": [\n              {\n                \"code\": \"mc_processing\",\n                \"amount\": {\n                  \"amount\": \"1\",\n                  \"currency\": \"USDC\"\n                },\n                \"asset\": {\n                  \"network\": \"ethereum_mainnet\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"description\": \"Illustrative bridge test commercial fee.\"\n              },\n              {\n                \"code\": \"bridge_protocol\",\n                \"amount\": {\n                  \"amount\": \"0.99\",\n                  \"currency\": \"USDC\"\n                },\n                \"asset\": {\n                  \"network\": \"arbitrum_one\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"description\": \"Illustrative bridge protocol cost at 100 basis points, explicitly accounted across the accepted bridge route.\"\n              }\n            ],\n            \"fee_treatment\": \"DEDUCT_FROM_DESTINATION\",\n            \"settlement_mode\": \"FIX_SOURCE\",\n            \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n            \"priced_at\": \"2026-10-23T10:00:00Z\",\n            \"expires_at\": \"2026-10-23T10:01:00Z\",\n            \"instructions\": {\n              \"source_account_ref\": \"acct_example_01\",\n              \"source_asset\": {\n                \"network\": \"ethereum_mainnet\",\n                \"currency\": \"USDC\",\n                \"kind\": \"TOKEN_CONTRACT\",\n                \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n              },\n              \"destination\": {\n                \"asset\": {\n                  \"network\": \"arbitrum_one\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"address\": \"0x2222222222222222222222222222222222222222\"\n              },\n              \"bridge\": {\n                \"route_ref\": \"illustrative_bridge_route_01\",\n                \"mode\": \"FAST\"\n              }\n            }\n          },\n          {\n            \"quote_ref\": \"bridge_fix_received_add_to_source\",\n            \"state\": \"EXECUTABLE\",\n            \"principal\": {\n              \"amount\": \"100\",\n              \"currency\": \"USDC\"\n            },\n            \"recipient_amount\": {\n              \"amount\": \"100\",\n              \"currency\": \"USDC\"\n            },\n            \"payer_total_debit\": {\n              \"amount\": \"102.010102\",\n              \"currency\": \"USDC\"\n            },\n            \"fees\": [\n              {\n                \"code\": \"mc_processing\",\n                \"amount\": {\n                  \"amount\": \"1\",\n                  \"currency\": \"USDC\"\n                },\n                \"asset\": {\n                  \"network\": \"ethereum_mainnet\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"description\": \"Illustrative bridge test commercial fee.\"\n              },\n              {\n                \"code\": \"bridge_protocol\",\n                \"amount\": {\n                  \"amount\": \"1.010102\",\n                  \"currency\": \"USDC\"\n                },\n                \"asset\": {\n                  \"network\": \"arbitrum_one\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"description\": \"Illustrative bridge protocol cost at 100 basis points, explicitly accounted across the accepted bridge route.\"\n              }\n            ],\n            \"fee_treatment\": \"ADD_TO_SOURCE\",\n            \"settlement_mode\": \"FIX_RECEIVED\",\n            \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n            \"priced_at\": \"2026-10-23T10:00:00Z\",\n            \"expires_at\": \"2026-10-23T10:01:00Z\",\n            \"instructions\": {\n              \"source_account_ref\": \"acct_example_01\",\n              \"source_asset\": {\n                \"network\": \"ethereum_mainnet\",\n                \"currency\": \"USDC\",\n                \"kind\": \"TOKEN_CONTRACT\",\n                \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n              },\n              \"destination\": {\n                \"asset\": {\n                  \"network\": \"arbitrum_one\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"address\": \"0x2222222222222222222222222222222222222222\"\n              },\n              \"bridge\": {\n                \"route_ref\": \"illustrative_bridge_route_01\",\n                \"mode\": \"FAST\"\n              }\n            }\n          },\n          {\n            \"quote_ref\": \"bridge_fix_received_deduct_from_destination\",\n            \"state\": \"EXECUTABLE\",\n            \"principal\": {\n              \"amount\": \"100\",\n              \"currency\": \"USDC\"\n            },\n            \"recipient_amount\": {\n              \"amount\": \"100\",\n              \"currency\": \"USDC\"\n            },\n            \"payer_total_debit\": {\n              \"amount\": \"102.010102\",\n              \"currency\": \"USDC\"\n            },\n            \"fees\": [\n              {\n                \"code\": \"mc_processing\",\n                \"amount\": {\n                  \"amount\": \"1\",\n                  \"currency\": \"USDC\"\n                },\n                \"asset\": {\n                  \"network\": \"ethereum_mainnet\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"description\": \"Illustrative bridge test commercial fee.\"\n              },\n              {\n                \"code\": \"bridge_protocol\",\n                \"amount\": {\n                  \"amount\": \"1.010102\",\n                  \"currency\": \"USDC\"\n                },\n                \"asset\": {\n                  \"network\": \"arbitrum_one\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"description\": \"Illustrative bridge protocol cost at 100 basis points, explicitly accounted across the accepted bridge route.\"\n              }\n            ],\n            \"fee_treatment\": \"DEDUCT_FROM_DESTINATION\",\n            \"settlement_mode\": \"FIX_RECEIVED\",\n            \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n            \"priced_at\": \"2026-10-23T10:00:00Z\",\n            \"expires_at\": \"2026-10-23T10:01:00Z\",\n            \"instructions\": {\n              \"source_account_ref\": \"acct_example_01\",\n              \"source_asset\": {\n                \"network\": \"ethereum_mainnet\",\n                \"currency\": \"USDC\",\n                \"kind\": \"TOKEN_CONTRACT\",\n                \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n              },\n              \"destination\": {\n                \"asset\": {\n                  \"network\": \"arbitrum_one\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"address\": \"0x2222222222222222222222222222222222222222\"\n              },\n              \"bridge\": {\n                \"route_ref\": \"illustrative_bridge_route_01\",\n                \"mode\": \"FAST\"\n              }\n            }\n          }\n        ],\n        \"x-example-source\": \"Sendico api/payments/quotation/internal/service/quotation/bridge_test.go: TestBridgeQuoteEconomicsAndExplicitPlan. Numeric examples adapted from fake gateway fixtures, not deployed-route or financial-accounting proof.\"\n      },\n      \"SubmitTransfer\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"quote_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"client_reference\": {\n            \"type\": \"string\",\n            \"description\": \"Caller reference persisted before submission. Does not substitute for Idempotency-Key.\"\n          }\n        },\n        \"required\": [\n          \"quote_ref\"\n        ],\n        \"description\": \"Maps to ExecutePayment. Tenant/actor and owner constraints are supplied by the trusted boundary. Same request/key must be replayed after a lost response. One public quote/intent admits at most one transfer across keys. A new payment requires a new quote. Never bypass expired identity reconciliation by replacing the key. Duplicate convergence returns the original client_reference.\"\n      },\n      \"ExecutionObservation\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"execution_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"network\": {\n            \"type\": \"string\"\n          },\n          \"status\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"CREATED\",\n              \"PROCESSING\",\n              \"WAITING_NETWORK\",\n              \"BROADCASTED\",\n              \"CONFIRMING\",\n              \"SUCCEEDED\",\n              \"FAILED\",\n              \"CANCELLED\"\n            ]\n          },\n          \"transaction_hash\": {\n            \"type\": \"string\",\n            \"description\": \"May change with an authoritative replacement transaction; not a financial effect identity.\"\n          }\n        },\n        \"required\": [\n          \"execution_ref\",\n          \"network\",\n          \"status\"\n        ],\n        \"description\": \"Read-only chain evidence. SUCCEEDED does not authorize a customer ledger posting.\"\n      },\n      \"BridgeStatus\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"bridge_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"stage\": {\n            \"type\": \"string\",\n            \"description\": \"Current stage of bridge processing, separate from financial settlement.\"\n          },\n          \"reason\": {\n            \"type\": \"string\"\n          },\n          \"approval_transaction_ref\": {\n            \"type\": \"string\"\n          },\n          \"burn_transaction_ref\": {\n            \"type\": \"string\"\n          },\n          \"mint_transaction_ref\": {\n            \"type\": \"string\"\n          }\n        },\n        \"required\": [\n          \"bridge_ref\",\n          \"stage\"\n        ],\n        \"description\": \"A burn establishes a protocol claim, not spendable destination custody. COMPLETED requires destination mint finality and Accounting receipts. Dependent payout completion is separate. Private wallet/funding/transit records are omitted.\"\n      },\n      \"Transfer\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"transfer_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"quote_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"client_reference\": {\n            \"type\": \"string\"\n          },\n          \"state\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"CREATED\",\n              \"EXECUTING\",\n              \"NEEDS_ATTENTION\",\n              \"SETTLED\",\n              \"FAILED\"\n            ],\n            \"x-allowed-transitions\": {\n              \"CREATED\": [\n                \"EXECUTING\",\n                \"FAILED\"\n              ],\n              \"EXECUTING\": [\n                \"NEEDS_ATTENTION\",\n                \"SETTLED\",\n                \"FAILED\"\n              ],\n              \"NEEDS_ATTENTION\": [\n                \"EXECUTING\",\n                \"SETTLED\",\n                \"FAILED\"\n              ],\n              \"SETTLED\": [],\n              \"FAILED\": []\n            },\n            \"description\": \"Existing aggregate transitions; SETTLED and FAILED are terminal for ordinary orchestration. Notifications/current reads can skip intermediate versions. Terminal does not prohibit new late-cost/corrective Accounting effects. See /guide.md.\"\n          },\n          \"version\": {\n            \"$ref\": \"#/components/schemas/Version\"\n          },\n          \"quote\": {\n            \"allOf\": [\n              {\n                \"$ref\": \"#/components/schemas/Quote\"\n              },\n              {\n                \"type\": \"object\",\n                \"properties\": {\n                  \"state\": {\n                    \"const\": \"EXECUTABLE\",\n                    \"type\": \"string\"\n                  }\n                },\n                \"required\": [\n                  \"state\"\n                ],\n                \"additionalProperties\": true\n              }\n            ],\n            \"description\": \"Immutable executable quote accepted by this transfer. It retains original expiry even when read after that time.\"\n          },\n          \"executions\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/ExecutionObservation\"\n            }\n          },\n          \"bridges\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/BridgeStatus\"\n            }\n          },\n          \"created_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"updated_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"completed_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"attention\": {\n            \"$ref\": \"#/components/schemas/Attention\"\n          }\n        },\n        \"required\": [\n          \"transfer_ref\",\n          \"quote_ref\",\n          \"state\",\n          \"version\",\n          \"quote\",\n          \"created_at\",\n          \"updated_at\"\n        ],\n        \"description\": \"An admitted digital asset transfer. SETTLED is the authoritative processing state; it is never inferred from a chain receipt. Financial receipts and account statements provide financial evidence. Each bridge lifecycle is independent. quote.instructions exposes the authoritative accepted source, destination, memo, and bridge choice. client_reference is caller correlation, never a deduplication identity or proof that the surrounding business payment in your application completed. attention is mandatory for NEEDS_ATTENTION and FAILED. WAIT describes automatic recovery without implying customer credit or immediate execution.\",\n        \"allOf\": [\n          {\n            \"if\": {\n              \"type\": \"object\",\n              \"properties\": {\n                \"state\": {\n                  \"enum\": [\n                    \"NEEDS_ATTENTION\",\n                    \"FAILED\"\n                  ]\n                }\n              },\n              \"required\": [\n                \"state\"\n              ],\n              \"additionalProperties\": true\n            },\n            \"then\": {\n              \"type\": \"object\",\n              \"required\": [\n                \"attention\"\n              ],\n              \"properties\": {\n                \"attention\": {\n                  \"$ref\": \"#/components/schemas/Attention\"\n                }\n              },\n              \"additionalProperties\": true\n            }\n          }\n        ],\n        \"examples\": [\n          {\n            \"transfer_ref\": \"transfer_example_01\",\n            \"quote_ref\": \"quote_example_01\",\n            \"client_reference\": \"merchant-withdrawal-2026-001\",\n            \"state\": \"CREATED\",\n            \"version\": \"1\",\n            \"quote\": {\n              \"quote_ref\": \"quote_example_01\",\n              \"state\": \"EXECUTABLE\",\n              \"principal\": {\n                \"amount\": \"125.50\",\n                \"currency\": \"USDC\"\n              },\n              \"recipient_amount\": {\n                \"amount\": \"125.50\",\n                \"currency\": \"USDC\"\n              },\n              \"payer_total_debit\": {\n                \"amount\": \"126.00\",\n                \"currency\": \"USDC\"\n              },\n              \"fees\": [\n                {\n                  \"code\": \"mc_processing\",\n                  \"amount\": {\n                    \"amount\": \"0.50\",\n                    \"currency\": \"USDC\"\n                  },\n                  \"description\": \"Illustrative Arctex processing fee\",\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  }\n                }\n              ],\n              \"fee_treatment\": \"ADD_TO_SOURCE\",\n              \"settlement_mode\": \"FIX_SOURCE\",\n              \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n              \"priced_at\": \"2026-10-23T10:00:00Z\",\n              \"expires_at\": \"2026-10-23T10:01:00Z\",\n              \"instructions\": {\n                \"source_account_ref\": \"acct_example_01\",\n                \"source_asset\": {\n                  \"network\": \"ethereum_sepolia\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"destination\": {\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  },\n                  \"address\": \"0x2222222222222222222222222222222222222222\"\n                }\n              }\n            },\n            \"executions\": [],\n            \"bridges\": [],\n            \"created_at\": \"2026-10-23T10:00:02Z\",\n            \"updated_at\": \"2026-10-23T10:00:02Z\"\n          },\n          {\n            \"transfer_ref\": \"transfer_example_01\",\n            \"quote_ref\": \"quote_example_01\",\n            \"client_reference\": \"merchant-withdrawal-2026-001\",\n            \"state\": \"FAILED\",\n            \"version\": \"3\",\n            \"quote\": {\n              \"quote_ref\": \"quote_example_01\",\n              \"state\": \"EXECUTABLE\",\n              \"principal\": {\n                \"amount\": \"125.50\",\n                \"currency\": \"USDC\"\n              },\n              \"recipient_amount\": {\n                \"amount\": \"125.50\",\n                \"currency\": \"USDC\"\n              },\n              \"payer_total_debit\": {\n                \"amount\": \"126.00\",\n                \"currency\": \"USDC\"\n              },\n              \"fees\": [\n                {\n                  \"code\": \"mc_processing\",\n                  \"amount\": {\n                    \"amount\": \"0.50\",\n                    \"currency\": \"USDC\"\n                  },\n                  \"description\": \"Illustrative Arctex processing fee\",\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  }\n                }\n              ],\n              \"fee_treatment\": \"ADD_TO_SOURCE\",\n              \"settlement_mode\": \"FIX_SOURCE\",\n              \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n              \"priced_at\": \"2026-10-23T10:00:00Z\",\n              \"expires_at\": \"2026-10-23T10:01:00Z\",\n              \"instructions\": {\n                \"source_account_ref\": \"acct_example_01\",\n                \"source_asset\": {\n                  \"network\": \"ethereum_sepolia\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"destination\": {\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  },\n                  \"address\": \"0x2222222222222222222222222222222222222222\"\n                }\n              }\n            },\n            \"executions\": [],\n            \"bridges\": [],\n            \"created_at\": \"2026-10-23T10:00:02Z\",\n            \"updated_at\": \"2026-10-23T10:01:00Z\",\n            \"completed_at\": \"2026-10-23T10:01:00Z\",\n            \"attention\": {\n              \"reason_code\": \"INSUFFICIENT_FUNDS\",\n              \"message\": \"Accounting reservation failed. No chain dispatch; any earlier completed hold must be recovered by the existing release workflow. Inspect committed effects.\",\n              \"action\": \"NONE\"\n            }\n          },\n          {\n            \"transfer_ref\": \"transfer_example_01\",\n            \"quote_ref\": \"quote_example_01\",\n            \"client_reference\": \"merchant-withdrawal-2026-001\",\n            \"state\": \"NEEDS_ATTENTION\",\n            \"version\": \"4\",\n            \"quote\": {\n              \"quote_ref\": \"quote_example_01\",\n              \"state\": \"EXECUTABLE\",\n              \"principal\": {\n                \"amount\": \"125.50\",\n                \"currency\": \"USDC\"\n              },\n              \"recipient_amount\": {\n                \"amount\": \"125.50\",\n                \"currency\": \"USDC\"\n              },\n              \"payer_total_debit\": {\n                \"amount\": \"126.00\",\n                \"currency\": \"USDC\"\n              },\n              \"fees\": [\n                {\n                  \"code\": \"mc_processing\",\n                  \"amount\": {\n                    \"amount\": \"0.50\",\n                    \"currency\": \"USDC\"\n                  },\n                  \"description\": \"Illustrative Arctex processing fee\",\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  }\n                }\n              ],\n              \"fee_treatment\": \"ADD_TO_SOURCE\",\n              \"settlement_mode\": \"FIX_SOURCE\",\n              \"readiness\": \"LIQUIDITY_OBTAINABLE\",\n              \"priced_at\": \"2026-10-23T10:00:00Z\",\n              \"expires_at\": \"2026-10-23T10:01:00Z\",\n              \"instructions\": {\n                \"source_account_ref\": \"acct_example_01\",\n                \"source_asset\": {\n                  \"network\": \"ethereum_sepolia\",\n                  \"currency\": \"USDC\",\n                  \"kind\": \"TOKEN_CONTRACT\",\n                  \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                },\n                \"destination\": {\n                  \"asset\": {\n                    \"network\": \"ethereum_sepolia\",\n                    \"currency\": \"USDC\",\n                    \"kind\": \"TOKEN_CONTRACT\",\n                    \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n                  },\n                  \"address\": \"0x2222222222222222222222222222222222222222\"\n                }\n              }\n            },\n            \"executions\": [],\n            \"bridges\": [],\n            \"created_at\": \"2026-10-23T10:00:02Z\",\n            \"updated_at\": \"2026-10-23T10:02:00Z\",\n            \"attention\": {\n              \"reason_code\": \"EXECUTION_UNCERTAIN\",\n              \"message\": \"Dispatch outcome is unresolved. Keep the original transfer identity and do not create a replacement payment.\",\n              \"action\": \"CONTACT_SUPPORT\"\n            }\n          }\n        ]\n      },\n      \"TransferList\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"items\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/Transfer\"\n            }\n          },\n          \"next_cursor\": {\n            \"type\": \"string\",\n            \"description\": \"Opaque next page token; omitted when exhausted.\"\n          }\n        },\n        \"required\": [\n          \"items\"\n        ]\n      },\n      \"BridgeRoute\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"route_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"source_asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          },\n          \"destination_asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          },\n          \"modes\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"type\": \"string\",\n              \"enum\": [\n                \"STANDARD\",\n                \"FAST\"\n              ]\n            }\n          }\n        },\n        \"required\": [\n          \"route_ref\",\n          \"source_asset\",\n          \"destination_asset\",\n          \"modes\"\n        ],\n        \"description\": \"Existing configured route/capability projection. Static eligibility does not promise current liquidity, signing readiness, or immediate dispatch. Obtain a concrete quote.\"\n      },\n      \"BridgeRouteList\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"items\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/BridgeRoute\"\n            }\n          }\n        },\n        \"required\": [\n          \"items\"\n        ]\n      },\n      \"Error\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"code\": {\n            \"type\": \"string\",\n            \"description\": \"Stable boundary error code.\",\n            \"examples\": [\n              \"IDEMPOTENCY_CONFLICT\"\n            ],\n            \"enum\": [\n              \"INVALID_REQUEST\",\n              \"UNAUTHENTICATED\",\n              \"PERMISSION_DENIED\",\n              \"NOT_FOUND\",\n              \"IDEMPOTENCY_CONFLICT\",\n              \"SUBMISSION_IN_PROGRESS\",\n              \"QUOTE_CLAIM_CONFLICT\",\n              \"QUOTE_EXPIRED\",\n              \"QUOTE_NOT_EXECUTABLE\",\n              \"UNSUPPORTED_QUOTE_COMBINATION\",\n              \"INVALID_INPUT\",\n              \"INVALID_PRECISION\",\n              \"INSUFFICIENT_FUNDS\",\n              \"LIQUIDITY_UNAVAILABLE\",\n              \"COMPLIANCE_REVIEW_REQUIRED\",\n              \"RECOVERY_CURSOR_EXPIRED\",\n              \"IDEMPOTENCY_RECORD_EXPIRED\",\n              \"RATE_LIMITED\",\n              \"DEPENDENCY_UNAVAILABLE\"\n            ]\n          },\n          \"message\": {\n            \"type\": \"string\",\n            \"description\": \"Human-readable explanation; do not branch on this text.\"\n          },\n          \"request_id\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          }\n        },\n        \"required\": [\n          \"code\",\n          \"message\",\n          \"request_id\"\n        ],\n        \"description\": \"A stable API error code with safe explanatory text. Handle the code rather than parsing message text. Retry guidance and request outcome certainty are in /guide.md. A lost response/503 never proves non-acceptance. Insufficient-funds or review failure discovered after acceptance appears in Transfer.attention under the existing transfer identity. The published /v1 code set is closed. New emitted codes require a new API major version. If a client nevertheless encounters an unknown code, preserve the request/key and response, treat acceptance as uncertain, and reconcile; never infer a new payment or refund. See /guide.md#v1-compatibility-and-unknown-values.\"\n      },\n      \"FinancialEffect\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"effect_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"journal_entry_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"posting_line_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"account_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"money\": {\n            \"$ref\": \"#/components/schemas/Money\"\n          },\n          \"committed_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"compensates_effect_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"origin\": {\n            \"$ref\": \"#/components/schemas/EffectOrigin\"\n          },\n          \"asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          },\n          \"line_type\": {\n            \"$ref\": \"#/components/schemas/PostingClassification\"\n          }\n        },\n        \"required\": [\n          \"effect_ref\",\n          \"journal_entry_ref\",\n          \"posting_line_ref\",\n          \"account_ref\",\n          \"money\",\n          \"committed_at\",\n          \"origin\",\n          \"asset\",\n          \"line_type\"\n        ],\n        \"description\": \"Immutable committed financial effect derived from authoritative Accounting receipt identities. effect_ref is stable for the same journal/posting line. It is not a transaction hash or delivery ID. A compensation is a new effect; its link is included only when supported by committed evidence. Same canonical object in GET, discovery, and financial_effect.committed. asset is the posting account liability identity and money.currency must match. origin is mandatory before publication; incomplete projection must stall/recover rather than emit an uncorrelated financial effect. Effect identity is never reused. Each principal/fee posting has its own effect; replay never charges or recognizes a fee again. line_type is the committed posting classification and is identical in the statement line, effect reads, and financial webhook. Classification does not change effect identity or authorize a customer posting in your application without its local business mapping. Effects can precede terminal transfer status (for example a posted hold, incurred fee, or compensation), and retained late costs/corrections can follow it. A reservation without a committed journal is not an effect. Recover by transfer_ref until the committed prefix is caught up and reconcile financial history; FAILED/SETTLED alone neither define refunds nor close future corrections. See /guide.md for outcome examples. committed_at is not an account-wide snapshot token. Neither journal identity nor money can be used to guess a source/cursor cutoff. The canonical effect contains no new synthetic account sequence.\",\n        \"examples\": [\n          {\n            \"effect_ref\": \"journal_example_01:line_example_01\",\n            \"journal_entry_ref\": \"journal_example_01\",\n            \"posting_line_ref\": \"line_example_01\",\n            \"account_ref\": \"acct_example_01\",\n            \"money\": {\n              \"amount\": \"-125.50\",\n              \"currency\": \"USDC\"\n            },\n            \"committed_at\": \"2026-10-23T10:02:00Z\",\n            \"origin\": {\n              \"type\": \"TRANSFER\",\n              \"transfer_ref\": \"transfer_example_01\",\n              \"client_reference\": \"merchant-withdrawal-2026-001\"\n            },\n            \"asset\": {\n              \"network\": \"ethereum_sepolia\",\n              \"currency\": \"USDC\",\n              \"kind\": \"TOKEN_CONTRACT\",\n              \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n            },\n            \"line_type\": \"MAIN\"\n          },\n          {\n            \"effect_ref\": \"journal_example_01:line_example_02\",\n            \"journal_entry_ref\": \"journal_example_01\",\n            \"posting_line_ref\": \"line_example_02\",\n            \"account_ref\": \"acct_example_01\",\n            \"money\": {\n              \"amount\": \"-0.50\",\n              \"currency\": \"USDC\"\n            },\n            \"committed_at\": \"2026-10-23T10:02:00Z\",\n            \"origin\": {\n              \"type\": \"TRANSFER\",\n              \"transfer_ref\": \"transfer_example_01\",\n              \"client_reference\": \"merchant-withdrawal-2026-001\"\n            },\n            \"asset\": {\n              \"network\": \"ethereum_sepolia\",\n              \"currency\": \"USDC\",\n              \"kind\": \"TOKEN_CONTRACT\",\n              \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n            },\n            \"line_type\": \"FEE\"\n          },\n          {\n            \"effect_ref\": \"journal_compensation_example_01:line_reverse_01\",\n            \"journal_entry_ref\": \"journal_compensation_example_01\",\n            \"posting_line_ref\": \"line_reverse_01\",\n            \"account_ref\": \"acct_example_01\",\n            \"money\": {\n              \"amount\": \"125.50\",\n              \"currency\": \"USDC\"\n            },\n            \"committed_at\": \"2026-10-24T10:00:00Z\",\n            \"origin\": {\n              \"type\": \"TRANSFER\",\n              \"transfer_ref\": \"transfer_example_01\",\n              \"client_reference\": \"merchant-withdrawal-2026-001\"\n            },\n            \"asset\": {\n              \"network\": \"ethereum_sepolia\",\n              \"currency\": \"USDC\",\n              \"kind\": \"TOKEN_CONTRACT\",\n              \"contract_address\": \"0x1111111111111111111111111111111111111111\"\n            },\n            \"line_type\": \"REVERSAL\",\n            \"compensates_effect_ref\": \"journal_example_01:line_example_01\"\n          }\n        ]\n      },\n      \"TransferEvent\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"event_id\": {\n            \"type\": \"string\",\n            \"description\": \"Stable durable business event ID; unchanged on redelivery.\"\n          },\n          \"delivery_id\": {\n            \"type\": \"string\",\n            \"description\": \"Unique delivery attempt ID, never a financial deduplication identity.\"\n          },\n          \"occurred_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"type\": {\n            \"const\": \"transfer.updated\",\n            \"type\": \"string\"\n          },\n          \"data\": {\n            \"$ref\": \"#/components/schemas/Transfer\"\n          }\n        },\n        \"required\": [\n          \"event_id\",\n          \"delivery_id\",\n          \"occurred_at\",\n          \"type\",\n          \"data\"\n        ],\n        \"description\": \"Resource lifecycle notification. Persist/compare the existing transfer version; never post customer financial entries from this event.\"\n      },\n      \"DepositEvent\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"event_id\": {\n            \"type\": \"string\",\n            \"description\": \"Stable durable business event ID; unchanged on redelivery.\"\n          },\n          \"delivery_id\": {\n            \"type\": \"string\",\n            \"description\": \"Unique delivery attempt ID, never a financial deduplication identity.\"\n          },\n          \"occurred_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"type\": {\n            \"const\": \"deposit.updated\",\n            \"type\": \"string\"\n          },\n          \"data\": {\n            \"$ref\": \"#/components/schemas/Deposit\"\n          }\n        },\n        \"required\": [\n          \"event_id\",\n          \"delivery_id\",\n          \"occurred_at\",\n          \"type\",\n          \"data\"\n        ],\n        \"description\": \"Deposit resource notification. Re-read the deposit if events arrive out of order. A financial posting is communicated separately from the committed Accounting receipt. Compare deposit_ref/version; recover missed or late revisions through /v1/deposits.\"\n      },\n      \"FinancialEvent\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"event_id\": {\n            \"type\": \"string\",\n            \"description\": \"Stable durable business event ID; unchanged on redelivery.\"\n          },\n          \"delivery_id\": {\n            \"type\": \"string\",\n            \"description\": \"Unique delivery attempt ID, never a financial deduplication identity.\"\n          },\n          \"occurred_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          },\n          \"type\": {\n            \"const\": \"financial_effect.committed\",\n            \"type\": \"string\"\n          },\n          \"data\": {\n            \"$ref\": \"#/components/schemas/FinancialEffect\"\n          }\n        },\n        \"required\": [\n          \"event_id\",\n          \"delivery_id\",\n          \"occurred_at\",\n          \"type\",\n          \"data\"\n        ],\n        \"description\": \"Originates exclusively from committed Accounting receipts/outbox. A financial consumer deduplicates effect_ref in the same transaction as its own customer posting. Typed origin and stable effect/posting identities are preserved by read recovery and by compensating entries.\"\n      },\n      \"ReceivingAttribution\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"receiving_address_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"route_version\": {\n            \"$ref\": \"#/components/schemas/Version\"\n          },\n          \"address\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          },\n          \"memo\": {\n            \"type\": \"string\"\n          },\n          \"captured_at\": {\n            \"$ref\": \"#/components/schemas/Timestamp\"\n          }\n        },\n        \"required\": [\n          \"address\",\n          \"captured_at\"\n        ],\n        \"description\": \"Receiving evidence retained with each deposit revision. Address, memo when present, and captured_at come from the original intake and remain immutable. receiving_address_ref and route_version are absent while authoritative route attribution is unresolved; never invent a reference or resolve historical ownership from the latest route. A later revision can add the proven route identity and version. Once established, that attribution remains unchanged; financial corrections append compensating effects. route_version requires receiving_address_ref and is included only when retained evidence proves it. Tenant visibility before route resolution must come from durable intake/custody ownership, never from caller-provided account or address filters. Your application retains its historical receiving-reference/address/memo-to-customer mapping.\",\n        \"allOf\": [\n          {\n            \"if\": {\n              \"required\": [\n                \"route_version\"\n              ],\n              \"type\": \"object\",\n              \"additionalProperties\": true,\n              \"properties\": {\n                \"route_version\": {\n                  \"$ref\": \"#/components/schemas/ReceivingAttribution/properties/route_version\"\n                }\n              }\n            },\n            \"then\": {\n              \"required\": [\n                \"receiving_address_ref\"\n              ],\n              \"type\": \"object\",\n              \"additionalProperties\": true,\n              \"properties\": {\n                \"receiving_address_ref\": {\n                  \"$ref\": \"#/components/schemas/ReceivingAttribution/properties/receiving_address_ref\"\n                }\n              }\n            }\n          }\n        ]\n      },\n      \"TransferInstructions\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"source_account_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"source_asset\": {\n            \"$ref\": \"#/components/schemas/Asset\"\n          },\n          \"destination\": {\n            \"$ref\": \"#/components/schemas/Destination\"\n          },\n          \"bridge\": {\n            \"$ref\": \"#/components/schemas/BridgeSelection\"\n          }\n        },\n        \"required\": [\n          \"source_account_ref\",\n          \"source_asset\",\n          \"destination\"\n        ],\n        \"description\": \"Immutable transfer instructions retained with the quote and accepted transfer, including a supplied destination memo and selected bridge. These instructions are not reconstructed from current routes or configuration. Accepted Transfer.quote.instructions remains unchanged after expiry, retries, or destination reconfiguration. Principal, amount mode, and fees are in the surrounding Quote.\"\n      },\n      \"DepositOrigin\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"type\": {\n            \"type\": \"string\",\n            \"const\": \"DEPOSIT\"\n          },\n          \"deposit_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          }\n        },\n        \"required\": [\n          \"type\",\n          \"deposit_ref\"\n        ],\n        \"description\": \"Origin of a committed deposit effect. Resolve receiving/customer attribution through the durable Deposit snapshot, not amount matching.\"\n      },\n      \"TransferOrigin\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"type\": {\n            \"type\": \"string\",\n            \"const\": \"TRANSFER\"\n          },\n          \"transfer_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          },\n          \"client_reference\": {\n            \"type\": \"string\"\n          }\n        },\n        \"required\": [\n          \"type\",\n          \"transfer_ref\"\n        ],\n        \"description\": \"Origin of a committed transfer effect, including its separately identified fee lines. Echo the persisted client reference when supplied; it is not the financial identity.\"\n      },\n      \"AdjustmentOrigin\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"type\": {\n            \"type\": \"string\",\n            \"const\": \"ADJUSTMENT\"\n          },\n          \"reason_code\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          },\n          \"case_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          }\n        },\n        \"required\": [\n          \"type\",\n          \"reason_code\"\n        ],\n        \"description\": \"Existing Accounting correction/reconciliation effect with a safe reason code and optional support case. This read projection does not introduce an adjustment API. Compensations link the original through FinancialEffect.compensates_effect_ref; never invent a deposit/transfer origin for an unrelated adjustment.\"\n      },\n      \"EffectOrigin\": {\n        \"oneOf\": [\n          {\n            \"$ref\": \"#/components/schemas/DepositOrigin\"\n          },\n          {\n            \"$ref\": \"#/components/schemas/TransferOrigin\"\n          },\n          {\n            \"$ref\": \"#/components/schemas/AdjustmentOrigin\"\n          }\n        ],\n        \"discriminator\": {\n          \"propertyName\": \"type\",\n          \"mapping\": {\n            \"DEPOSIT\": \"#/components/schemas/DepositOrigin\",\n            \"TRANSFER\": \"#/components/schemas/TransferOrigin\",\n            \"ADJUSTMENT\": \"#/components/schemas/AdjustmentOrigin\"\n          }\n        },\n        \"description\": \"Structured business origin persisted with the committed receipt or durably joined by its stable identity before external publication. No matching by hash, amount, or description.\"\n      },\n      \"Attention\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"reason_code\": {\n            \"type\": \"string\",\n            \"minLength\": 1,\n            \"enum\": [\n              \"ROUTE_UNRESOLVED\",\n              \"COMPLIANCE_REVIEW_REQUIRED\",\n              \"DEPOSIT_REJECTED\",\n              \"SWEEP_RETRY_PENDING\",\n              \"CREDIT_RETRY_PENDING\",\n              \"REMEDIATION_IN_PROGRESS\",\n              \"INSUFFICIENT_FUNDS\",\n              \"LIQUIDITY_UNAVAILABLE\",\n              \"EXECUTION_UNCERTAIN\",\n              \"ACCOUNTING_SETTLEMENT_PENDING\",\n              \"COMPENSATION_PENDING\",\n              \"EXECUTION_FAILED\",\n              \"OPERATOR_REVIEW_REQUIRED\"\n            ]\n          },\n          \"message\": {\n            \"type\": \"string\"\n          },\n          \"action\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"WAIT\",\n              \"CONTACT_SUPPORT\",\n              \"NONE\"\n            ]\n          },\n          \"case_ref\": {\n            \"$ref\": \"#/components/schemas/Reference\"\n          }\n        },\n        \"required\": [\n          \"reason_code\",\n          \"message\",\n          \"action\"\n        ],\n        \"description\": \"Tenant-safe explanation and action guidance for review, rejection, or operational attention. Stable reason codes provide guidance without disclosing private compliance evidence or internal diagnostics. Arctex operations handles digital asset review and return execution; your application retains its customer and business decisions. No new public approval, return, or cancellation path. The catalog/action guidance is in /guide.md. Unknown private reasons map to OPERATOR_REVIEW_REQUIRED rather than leaking internals. attention is mandatory for the resource states constrained below. It does not itself prove a reservation release or committed financial effect. Published reason/action values are frozen within /v1. Unknown received values must retain the raw record and prevent dependent automated financial decisions; do not map them to NONE/success.\"\n      },\n      \"DepositList\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"items\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/Deposit\"\n            }\n          },\n          \"next_cursor\": {\n            \"type\": \"string\",\n            \"minLength\": 1,\n            \"description\": \"Durable opaque resume checkpoint, always returned, even for empty/caught-up pages.\"\n          },\n          \"has_more\": {\n            \"type\": \"boolean\",\n            \"description\": \"More items currently available after this page; false does not close the stream.\"\n          }\n        },\n        \"required\": [\n          \"items\",\n          \"next_cursor\",\n          \"has_more\"\n        ],\n        \"description\": \"Omit cursor to read retained history. Continue using next_cursor, including after has_more=false or an empty page. Persist a returned cursor only after durably handling every item. Tokens bind the authenticated tenant, this surface, and unchanged filters; malformed/mismatched tokens return 400, expired or unrecoverable checkpoints return 410 and require reconciliation. Never silently reset. Cursor order is durable projection commit order, not created_at/event_time, so late commits and state updates remain discoverable. Projection recovery must not advance past unprojected committed source changes. Current recovery cursors expire after 24 hours; retained revisions/effects have no automatic TTL. Production retention, archival and replay agreements remain release gates; see /guide.md#installed-retention-and-delivery-policy. Each item is an immutable retained Deposit revision. The same deposit_ref can recur at a later version, including a late credit; deduplicate/compare by deposit_ref and version. GET by reference returns the current revision. Version need not be contiguous. No generic event API or new deposit lifecycle is introduced.\"\n      },\n      \"FinancialEffectList\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"items\": {\n            \"type\": \"array\",\n            \"items\": {\n              \"$ref\": \"#/components/schemas/FinancialEffect\"\n            }\n          },\n          \"next_cursor\": {\n            \"type\": \"string\",\n            \"minLength\": 1,\n            \"description\": \"Durable opaque resume checkpoint, always returned, even for empty/caught-up pages.\"\n          },\n          \"has_more\": {\n            \"type\": \"boolean\",\n            \"description\": \"More items currently available after this page. false does not close the stream, prove source Accounting is fully projected, or align this page with an independently read account balance.\"\n          }\n        },\n        \"required\": [\n          \"items\",\n          \"next_cursor\",\n          \"has_more\"\n        ],\n        \"description\": \"Omit cursor to read retained history. Continue using next_cursor, including after has_more=false or an empty page. Persist a returned cursor only after durably handling every item. Tokens bind the authenticated tenant, this surface, and unchanged filters; malformed/mismatched tokens return 400, expired or unrecoverable checkpoints return 410 and require reconciliation. Never silently reset. Cursor order is durable projection commit order, not created_at/event_time, so late commits and state updates remain discoverable. Projection recovery must not advance past unprojected committed source changes. Current recovery cursors expire after 24 hours; retained revisions/effects have no automatic TTL. Production retention, archival and replay agreements remain release gates; see /guide.md#installed-retention-and-delivery-policy. Each item is an immutable committed effect. The same effect_ref always has the same financial content, whether read here, by reference, or delivered by webhook. Neither late event_time nor a failed webhook removes it from recovery. This projection checkpoint has no published numeric relationship to Balance.version. For a financial proof use the Accounting snapshot/export procedure and exact source posting membership described in /guide.md#consistent-reconciliation-cutoff. Catching up to the current projection head alone cannot detect an unprojected source commit.\"\n      },\n      \"PostingClassification\": {\n        \"type\": \"string\",\n        \"enum\": [\n          \"MAIN\",\n          \"FEE\",\n          \"SPREAD\",\n          \"REVERSAL\"\n        ],\n        \"description\": \"Existing committed Ledger LineType projection. Preserve the receipt classification; never infer principal, fee, spread, or reversal from amount, sign, effect ordering, or a description. MAIN includes primary debit/credit lines. A reversal may retain another line classification under an existing posting plan, so compensation is identified by its committed linkage rather than this enum alone. Unsupported or unspecified source classifications block publication until explicitly mapped; they are not silently coerced to MAIN.\"\n      },\n      \"ChainEvidence\": {\n        \"type\": \"object\",\n        \"additionalProperties\": false,\n        \"properties\": {\n          \"network\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          },\n          \"transaction_hash\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          },\n          \"transfer_identity\": {\n            \"type\": \"string\",\n            \"minLength\": 1,\n            \"description\": \"Optional opaque canonical transfer locator from the chain observer. EVM token logs bind transaction hash + log index; TRON events/transfers bind hash + the observer event/transfer position; native transfers bind the transaction and trace/output position when multiple transfers exist. Network is always part of identity. No positional value is invented when not retained. Never use this instead of deposit_ref/effect_ref.\"\n          },\n          \"finality\": {\n            \"type\": \"string\",\n            \"enum\": [\n              \"OBSERVED\",\n              \"CONFIRMING\",\n              \"FINALIZED\",\n              \"REORGED\"\n            ],\n            \"description\": \"Optional authoritative chain observation, omitted if unavailable. FINALIZED reflects the configured network finality policy, not an absolute guarantee or a financial credit.\"\n          },\n          \"block_hash\": {\n            \"type\": \"string\",\n            \"minLength\": 1\n          },\n          \"block_height\": {\n            \"$ref\": \"#/components/schemas/Version\"\n          }\n        },\n        \"required\": [\n          \"network\",\n          \"transaction_hash\"\n        ],\n        \"description\": \"Tenant-safe evidence from persisted DepositObservation.TransactionHash and canonical observer receipts, when retained. Unavailable locator/finality fields are omitted; they are not implied by deposit state. Replacement/reorg updates evidence in later deposit revisions; preserve recognition identity for the same economic transfer and use append-only financial corrections after committed credit. A distinct on-chain transfer receives its own producer identity. No new chain indexing/replacement engine is introduced.\"\n      }\n    }\n  },\n  \"x-review-source-revision\": \"c9d7dd5b10789dad90017aa14fd5013aad1970b0\",\n  \"x-compatibility-policy\": {\n    \"document_version\": \"Draft document version is independent of /v1. Pre-release revisions may change the draft explicitly; these are not production compatibility claims.\",\n    \"request_comparison\": \"stored-validated-json-values-and-presence; exact decimal strings; frozen catalog resolution\",\n    \"response_shapes\": \"Closed published /v1 shapes, response discriminators/enums and error/attention code sets are frozen; additions to existing closed responses are breaking.\",\n    \"cutoff\": \"No public balance-version/effect-cursor equivalence; use Accounting-backed reconciliation procedure.\",\n    \"generated_client_gate\": \"Implemented: canonical local generator, generated Go client and portable consumer guards; make integration-export verifies isolated offline tests/vet, fingerprints, byte-identical exports and exact regeneration. Client integration acceptance remains separate.\"\n  },\n  \"x-validation-date\": \"2026-10-26\"\n}\n"

type Account struct {
	AccountCode      *string                    `json:"account_code,omitempty"`
	AccountRef       Reference                  `json:"account_ref"`
	Asset            Asset                      `json:"asset"`
	Status           string                     `json:"status"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *Account) UnmarshalJSON(raw []byte) error {
	type plain Account
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "account_code")
	delete(extra, "account_ref")
	delete(extra, "asset")
	delete(extra, "status")
	p.AdditionalFields = extra
	*v = Account(p)
	return nil
}
func (v Account) MarshalJSON() ([]byte, error) {
	type plain Account
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "account_code":
			continue
		case "account_ref":
			continue
		case "asset":
			continue
		case "status":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type AccountList struct {
	Items            []Account                  `json:"items"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *AccountList) UnmarshalJSON(raw []byte) error {
	type plain AccountList
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "items")
	p.AdditionalFields = extra
	*v = AccountList(p)
	return nil
}
func (v AccountList) MarshalJSON() ([]byte, error) {
	type plain AccountList
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "items":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type AdjustmentOrigin struct {
	CaseRef          *Reference                 `json:"case_ref,omitempty"`
	ReasonCode       string                     `json:"reason_code"`
	Type             string                     `json:"type"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *AdjustmentOrigin) UnmarshalJSON(raw []byte) error {
	type plain AdjustmentOrigin
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "case_ref")
	delete(extra, "reason_code")
	delete(extra, "type")
	p.AdditionalFields = extra
	*v = AdjustmentOrigin(p)
	return nil
}
func (v AdjustmentOrigin) MarshalJSON() ([]byte, error) {
	type plain AdjustmentOrigin
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "case_ref":
			continue
		case "reason_code":
			continue
		case "type":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type Asset struct {
	ContractAddress  *string                    `json:"contract_address,omitempty"`
	Currency         string                     `json:"currency"`
	Kind             string                     `json:"kind"`
	Network          string                     `json:"network"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *Asset) UnmarshalJSON(raw []byte) error {
	type plain Asset
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "contract_address")
	delete(extra, "currency")
	delete(extra, "kind")
	delete(extra, "network")
	p.AdditionalFields = extra
	*v = Asset(p)
	return nil
}
func (v Asset) MarshalJSON() ([]byte, error) {
	type plain Asset
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "contract_address":
			continue
		case "currency":
			continue
		case "kind":
			continue
		case "network":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type AssetCatalogEntry struct {
	AccountingScale  int64                      `json:"accounting_scale"`
	Asset            Asset                      `json:"asset"`
	DisplayScale     int64                      `json:"display_scale"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *AssetCatalogEntry) UnmarshalJSON(raw []byte) error {
	type plain AssetCatalogEntry
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "accounting_scale")
	delete(extra, "asset")
	delete(extra, "display_scale")
	p.AdditionalFields = extra
	*v = AssetCatalogEntry(p)
	return nil
}
func (v AssetCatalogEntry) MarshalJSON() ([]byte, error) {
	type plain AssetCatalogEntry
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "accounting_scale":
			continue
		case "asset":
			continue
		case "display_scale":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type AssetList struct {
	CatalogRevision  string                     `json:"catalog_revision"`
	Items            []AssetCatalogEntry        `json:"items"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *AssetList) UnmarshalJSON(raw []byte) error {
	type plain AssetList
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "catalog_revision")
	delete(extra, "items")
	p.AdditionalFields = extra
	*v = AssetList(p)
	return nil
}
func (v AssetList) MarshalJSON() ([]byte, error) {
	type plain AssetList
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "catalog_revision":
			continue
		case "items":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type Attention struct {
	Action           string                     `json:"action"`
	CaseRef          *Reference                 `json:"case_ref,omitempty"`
	Message          string                     `json:"message"`
	ReasonCode       string                     `json:"reason_code"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *Attention) UnmarshalJSON(raw []byte) error {
	type plain Attention
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "action")
	delete(extra, "case_ref")
	delete(extra, "message")
	delete(extra, "reason_code")
	p.AdditionalFields = extra
	*v = Attention(p)
	return nil
}
func (v Attention) MarshalJSON() ([]byte, error) {
	type plain Attention
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "action":
			continue
		case "case_ref":
			continue
		case "message":
			continue
		case "reason_code":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type Balance struct {
	AccountRef       Reference                  `json:"account_ref"`
	AsOf             *Timestamp                 `json:"as_of,omitempty"`
	Asset            Asset                      `json:"asset"`
	Balance          Money                      `json:"balance"`
	LastUpdated      *Timestamp                 `json:"last_updated,omitempty"`
	Version          *Version                   `json:"version,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *Balance) UnmarshalJSON(raw []byte) error {
	type plain Balance
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "account_ref")
	delete(extra, "as_of")
	delete(extra, "asset")
	delete(extra, "balance")
	delete(extra, "last_updated")
	delete(extra, "version")
	p.AdditionalFields = extra
	*v = Balance(p)
	return nil
}
func (v Balance) MarshalJSON() ([]byte, error) {
	type plain Balance
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "account_ref":
			continue
		case "as_of":
			continue
		case "asset":
			continue
		case "balance":
			continue
		case "last_updated":
			continue
		case "version":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type BridgeRoute struct {
	DestinationAsset Asset                      `json:"destination_asset"`
	Modes            []string                   `json:"modes"`
	RouteRef         Reference                  `json:"route_ref"`
	SourceAsset      Asset                      `json:"source_asset"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *BridgeRoute) UnmarshalJSON(raw []byte) error {
	type plain BridgeRoute
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "destination_asset")
	delete(extra, "modes")
	delete(extra, "route_ref")
	delete(extra, "source_asset")
	p.AdditionalFields = extra
	*v = BridgeRoute(p)
	return nil
}
func (v BridgeRoute) MarshalJSON() ([]byte, error) {
	type plain BridgeRoute
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "destination_asset":
			continue
		case "modes":
			continue
		case "route_ref":
			continue
		case "source_asset":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type BridgeRouteList struct {
	Items            []BridgeRoute              `json:"items"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *BridgeRouteList) UnmarshalJSON(raw []byte) error {
	type plain BridgeRouteList
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "items")
	p.AdditionalFields = extra
	*v = BridgeRouteList(p)
	return nil
}
func (v BridgeRouteList) MarshalJSON() ([]byte, error) {
	type plain BridgeRouteList
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "items":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type BridgeSelection struct {
	Mode             string                     `json:"mode"`
	RouteRef         Reference                  `json:"route_ref"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *BridgeSelection) UnmarshalJSON(raw []byte) error {
	type plain BridgeSelection
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "mode")
	delete(extra, "route_ref")
	p.AdditionalFields = extra
	*v = BridgeSelection(p)
	return nil
}
func (v BridgeSelection) MarshalJSON() ([]byte, error) {
	type plain BridgeSelection
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "mode":
			continue
		case "route_ref":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type BridgeStatus struct {
	ApprovalTransactionRef *string                    `json:"approval_transaction_ref,omitempty"`
	BridgeRef              Reference                  `json:"bridge_ref"`
	BurnTransactionRef     *string                    `json:"burn_transaction_ref,omitempty"`
	MintTransactionRef     *string                    `json:"mint_transaction_ref,omitempty"`
	Reason                 *string                    `json:"reason,omitempty"`
	Stage                  string                     `json:"stage"`
	AdditionalFields       map[string]json.RawMessage `json:"-"`
}

func (v *BridgeStatus) UnmarshalJSON(raw []byte) error {
	type plain BridgeStatus
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "approval_transaction_ref")
	delete(extra, "bridge_ref")
	delete(extra, "burn_transaction_ref")
	delete(extra, "mint_transaction_ref")
	delete(extra, "reason")
	delete(extra, "stage")
	p.AdditionalFields = extra
	*v = BridgeStatus(p)
	return nil
}
func (v BridgeStatus) MarshalJSON() ([]byte, error) {
	type plain BridgeStatus
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "approval_transaction_ref":
			continue
		case "bridge_ref":
			continue
		case "burn_transaction_ref":
			continue
		case "mint_transaction_ref":
			continue
		case "reason":
			continue
		case "stage":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type ChainEvidence struct {
	BlockHash        *string                    `json:"block_hash,omitempty"`
	BlockHeight      *Version                   `json:"block_height,omitempty"`
	Finality         *string                    `json:"finality,omitempty"`
	Network          string                     `json:"network"`
	TransactionHash  string                     `json:"transaction_hash"`
	TransferIdentity *string                    `json:"transfer_identity,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *ChainEvidence) UnmarshalJSON(raw []byte) error {
	type plain ChainEvidence
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "block_hash")
	delete(extra, "block_height")
	delete(extra, "finality")
	delete(extra, "network")
	delete(extra, "transaction_hash")
	delete(extra, "transfer_identity")
	p.AdditionalFields = extra
	*v = ChainEvidence(p)
	return nil
}
func (v ChainEvidence) MarshalJSON() ([]byte, error) {
	type plain ChainEvidence
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "block_hash":
			continue
		case "block_height":
			continue
		case "finality":
			continue
		case "network":
			continue
		case "transaction_hash":
			continue
		case "transfer_identity":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type CreateReceivingAddress struct {
	AccountRef       Reference                  `json:"account_ref"`
	Asset            Asset                      `json:"asset"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *CreateReceivingAddress) UnmarshalJSON(raw []byte) error {
	type plain CreateReceivingAddress
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "account_ref")
	delete(extra, "asset")
	p.AdditionalFields = extra
	*v = CreateReceivingAddress(p)
	return nil
}
func (v CreateReceivingAddress) MarshalJSON() ([]byte, error) {
	type plain CreateReceivingAddress
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "account_ref":
			continue
		case "asset":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type Decimal = string
type Deposit struct {
	AccountRef            *Reference                 `json:"account_ref,omitempty"`
	Amount                PositiveDecimal            `json:"amount"`
	Asset                 Asset                      `json:"asset"`
	Attention             *Attention                 `json:"attention,omitempty"`
	ChainEvidence         *ChainEvidence             `json:"chain_evidence,omitempty"`
	CreatedAt             Timestamp                  `json:"created_at"`
	CreditJournalEntryRef *Reference                 `json:"credit_journal_entry_ref,omitempty"`
	DepositRef            Reference                  `json:"deposit_ref"`
	InboundRef            Reference                  `json:"inbound_ref"`
	Receiving             ReceivingAttribution       `json:"receiving"`
	State                 string                     `json:"state"`
	UpdatedAt             Timestamp                  `json:"updated_at"`
	Version               Version                    `json:"version"`
	AdditionalFields      map[string]json.RawMessage `json:"-"`
}

func (v *Deposit) UnmarshalJSON(raw []byte) error {
	type plain Deposit
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "account_ref")
	delete(extra, "amount")
	delete(extra, "asset")
	delete(extra, "attention")
	delete(extra, "chain_evidence")
	delete(extra, "created_at")
	delete(extra, "credit_journal_entry_ref")
	delete(extra, "deposit_ref")
	delete(extra, "inbound_ref")
	delete(extra, "receiving")
	delete(extra, "state")
	delete(extra, "updated_at")
	delete(extra, "version")
	p.AdditionalFields = extra
	*v = Deposit(p)
	return nil
}
func (v Deposit) MarshalJSON() ([]byte, error) {
	type plain Deposit
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "account_ref":
			continue
		case "amount":
			continue
		case "asset":
			continue
		case "attention":
			continue
		case "chain_evidence":
			continue
		case "created_at":
			continue
		case "credit_journal_entry_ref":
			continue
		case "deposit_ref":
			continue
		case "inbound_ref":
			continue
		case "receiving":
			continue
		case "state":
			continue
		case "updated_at":
			continue
		case "version":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type DepositEvent struct {
	Data             Deposit                    `json:"data"`
	DeliveryId       string                     `json:"delivery_id"`
	EventId          string                     `json:"event_id"`
	OccurredAt       Timestamp                  `json:"occurred_at"`
	Type             string                     `json:"type"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *DepositEvent) UnmarshalJSON(raw []byte) error {
	type plain DepositEvent
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "data")
	delete(extra, "delivery_id")
	delete(extra, "event_id")
	delete(extra, "occurred_at")
	delete(extra, "type")
	p.AdditionalFields = extra
	*v = DepositEvent(p)
	return nil
}
func (v DepositEvent) MarshalJSON() ([]byte, error) {
	type plain DepositEvent
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "data":
			continue
		case "delivery_id":
			continue
		case "event_id":
			continue
		case "occurred_at":
			continue
		case "type":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type DepositList struct {
	HasMore          bool                       `json:"has_more"`
	Items            []Deposit                  `json:"items"`
	NextCursor       string                     `json:"next_cursor"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *DepositList) UnmarshalJSON(raw []byte) error {
	type plain DepositList
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "has_more")
	delete(extra, "items")
	delete(extra, "next_cursor")
	p.AdditionalFields = extra
	*v = DepositList(p)
	return nil
}
func (v DepositList) MarshalJSON() ([]byte, error) {
	type plain DepositList
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "has_more":
			continue
		case "items":
			continue
		case "next_cursor":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type DepositOrigin struct {
	DepositRef       Reference                  `json:"deposit_ref"`
	Type             string                     `json:"type"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *DepositOrigin) UnmarshalJSON(raw []byte) error {
	type plain DepositOrigin
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "deposit_ref")
	delete(extra, "type")
	p.AdditionalFields = extra
	*v = DepositOrigin(p)
	return nil
}
func (v DepositOrigin) MarshalJSON() ([]byte, error) {
	type plain DepositOrigin
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "deposit_ref":
			continue
		case "type":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type Destination struct {
	Address          string                     `json:"address"`
	Asset            Asset                      `json:"asset"`
	Memo             *string                    `json:"memo,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *Destination) UnmarshalJSON(raw []byte) error {
	type plain Destination
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "address")
	delete(extra, "asset")
	delete(extra, "memo")
	p.AdditionalFields = extra
	*v = Destination(p)
	return nil
}
func (v Destination) MarshalJSON() ([]byte, error) {
	type plain Destination
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "address":
			continue
		case "asset":
			continue
		case "memo":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type EffectOrigin = json.RawMessage
type Error struct {
	Code             string                     `json:"code"`
	Message          string                     `json:"message"`
	RequestId        Reference                  `json:"request_id"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *Error) UnmarshalJSON(raw []byte) error {
	type plain Error
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "code")
	delete(extra, "message")
	delete(extra, "request_id")
	p.AdditionalFields = extra
	*v = Error(p)
	return nil
}
func (v Error) MarshalJSON() ([]byte, error) {
	type plain Error
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "code":
			continue
		case "message":
			continue
		case "request_id":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type ExecutionObservation struct {
	ExecutionRef     Reference                  `json:"execution_ref"`
	Network          string                     `json:"network"`
	Status           string                     `json:"status"`
	TransactionHash  *string                    `json:"transaction_hash,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *ExecutionObservation) UnmarshalJSON(raw []byte) error {
	type plain ExecutionObservation
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "execution_ref")
	delete(extra, "network")
	delete(extra, "status")
	delete(extra, "transaction_hash")
	p.AdditionalFields = extra
	*v = ExecutionObservation(p)
	return nil
}
func (v ExecutionObservation) MarshalJSON() ([]byte, error) {
	type plain ExecutionObservation
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "execution_ref":
			continue
		case "network":
			continue
		case "status":
			continue
		case "transaction_hash":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type Fee struct {
	Amount           Money                      `json:"amount"`
	Asset            Asset                      `json:"asset"`
	Code             string                     `json:"code"`
	Description      *string                    `json:"description,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *Fee) UnmarshalJSON(raw []byte) error {
	type plain Fee
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "amount")
	delete(extra, "asset")
	delete(extra, "code")
	delete(extra, "description")
	p.AdditionalFields = extra
	*v = Fee(p)
	return nil
}
func (v Fee) MarshalJSON() ([]byte, error) {
	type plain Fee
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "amount":
			continue
		case "asset":
			continue
		case "code":
			continue
		case "description":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type FinancialEffect struct {
	AccountRef           Reference                  `json:"account_ref"`
	Asset                Asset                      `json:"asset"`
	CommittedAt          Timestamp                  `json:"committed_at"`
	CompensatesEffectRef *Reference                 `json:"compensates_effect_ref,omitempty"`
	EffectRef            Reference                  `json:"effect_ref"`
	JournalEntryRef      Reference                  `json:"journal_entry_ref"`
	LineType             PostingClassification      `json:"line_type"`
	Money                Money                      `json:"money"`
	Origin               EffectOrigin               `json:"origin"`
	PostingLineRef       Reference                  `json:"posting_line_ref"`
	AdditionalFields     map[string]json.RawMessage `json:"-"`
}

func (v *FinancialEffect) UnmarshalJSON(raw []byte) error {
	type plain FinancialEffect
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "account_ref")
	delete(extra, "asset")
	delete(extra, "committed_at")
	delete(extra, "compensates_effect_ref")
	delete(extra, "effect_ref")
	delete(extra, "journal_entry_ref")
	delete(extra, "line_type")
	delete(extra, "money")
	delete(extra, "origin")
	delete(extra, "posting_line_ref")
	p.AdditionalFields = extra
	*v = FinancialEffect(p)
	return nil
}
func (v FinancialEffect) MarshalJSON() ([]byte, error) {
	type plain FinancialEffect
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "account_ref":
			continue
		case "asset":
			continue
		case "committed_at":
			continue
		case "compensates_effect_ref":
			continue
		case "effect_ref":
			continue
		case "journal_entry_ref":
			continue
		case "line_type":
			continue
		case "money":
			continue
		case "origin":
			continue
		case "posting_line_ref":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type FinancialEffectList struct {
	HasMore          bool                       `json:"has_more"`
	Items            []FinancialEffect          `json:"items"`
	NextCursor       string                     `json:"next_cursor"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *FinancialEffectList) UnmarshalJSON(raw []byte) error {
	type plain FinancialEffectList
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "has_more")
	delete(extra, "items")
	delete(extra, "next_cursor")
	p.AdditionalFields = extra
	*v = FinancialEffectList(p)
	return nil
}
func (v FinancialEffectList) MarshalJSON() ([]byte, error) {
	type plain FinancialEffectList
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "has_more":
			continue
		case "items":
			continue
		case "next_cursor":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type FinancialEvent struct {
	Data             FinancialEffect            `json:"data"`
	DeliveryId       string                     `json:"delivery_id"`
	EventId          string                     `json:"event_id"`
	OccurredAt       Timestamp                  `json:"occurred_at"`
	Type             string                     `json:"type"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *FinancialEvent) UnmarshalJSON(raw []byte) error {
	type plain FinancialEvent
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "data")
	delete(extra, "delivery_id")
	delete(extra, "event_id")
	delete(extra, "occurred_at")
	delete(extra, "type")
	p.AdditionalFields = extra
	*v = FinancialEvent(p)
	return nil
}
func (v FinancialEvent) MarshalJSON() ([]byte, error) {
	type plain FinancialEvent
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "data":
			continue
		case "delivery_id":
			continue
		case "event_id":
			continue
		case "occurred_at":
			continue
		case "type":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type Money struct {
	Amount           Decimal                    `json:"amount"`
	Currency         string                     `json:"currency"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *Money) UnmarshalJSON(raw []byte) error {
	type plain Money
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "amount")
	delete(extra, "currency")
	p.AdditionalFields = extra
	*v = Money(p)
	return nil
}
func (v Money) MarshalJSON() ([]byte, error) {
	type plain Money
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "amount":
			continue
		case "currency":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type PositiveDecimal = string
type PostingClassification = string
type PostingLine struct {
	AccountRef       Reference                  `json:"account_ref"`
	Asset            Asset                      `json:"asset"`
	EffectRef        Reference                  `json:"effect_ref"`
	LineType         PostingClassification      `json:"line_type"`
	Money            Money                      `json:"money"`
	Origin           EffectOrigin               `json:"origin"`
	PostingLineRef   Reference                  `json:"posting_line_ref"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *PostingLine) UnmarshalJSON(raw []byte) error {
	type plain PostingLine
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "account_ref")
	delete(extra, "asset")
	delete(extra, "effect_ref")
	delete(extra, "line_type")
	delete(extra, "money")
	delete(extra, "origin")
	delete(extra, "posting_line_ref")
	p.AdditionalFields = extra
	*v = PostingLine(p)
	return nil
}
func (v PostingLine) MarshalJSON() ([]byte, error) {
	type plain PostingLine
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "account_ref":
			continue
		case "asset":
			continue
		case "effect_ref":
			continue
		case "line_type":
			continue
		case "money":
			continue
		case "origin":
			continue
		case "posting_line_ref":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type Quote struct {
	BlockReason      *string                    `json:"block_reason,omitempty"`
	ExpiresAt        *Timestamp                 `json:"expires_at,omitempty"`
	FeeTreatment     string                     `json:"fee_treatment"`
	Fees             *[]Fee                     `json:"fees,omitempty"`
	Instructions     TransferInstructions       `json:"instructions"`
	PayerTotalDebit  *Money                     `json:"payer_total_debit,omitempty"`
	PricedAt         *Timestamp                 `json:"priced_at,omitempty"`
	Principal        *Money                     `json:"principal,omitempty"`
	QuoteRef         Reference                  `json:"quote_ref"`
	Readiness        string                     `json:"readiness"`
	RecipientAmount  *Money                     `json:"recipient_amount,omitempty"`
	SettlementMode   string                     `json:"settlement_mode"`
	State            string                     `json:"state"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *Quote) UnmarshalJSON(raw []byte) error {
	type plain Quote
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "block_reason")
	delete(extra, "expires_at")
	delete(extra, "fee_treatment")
	delete(extra, "fees")
	delete(extra, "instructions")
	delete(extra, "payer_total_debit")
	delete(extra, "priced_at")
	delete(extra, "principal")
	delete(extra, "quote_ref")
	delete(extra, "readiness")
	delete(extra, "recipient_amount")
	delete(extra, "settlement_mode")
	delete(extra, "state")
	p.AdditionalFields = extra
	*v = Quote(p)
	return nil
}
func (v Quote) MarshalJSON() ([]byte, error) {
	type plain Quote
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "block_reason":
			continue
		case "expires_at":
			continue
		case "fee_treatment":
			continue
		case "fees":
			continue
		case "instructions":
			continue
		case "payer_total_debit":
			continue
		case "priced_at":
			continue
		case "principal":
			continue
		case "quote_ref":
			continue
		case "readiness":
			continue
		case "recipient_amount":
			continue
		case "settlement_mode":
			continue
		case "state":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type QuoteRequest struct {
	Amount           PositiveDecimal            `json:"amount"`
	Bridge           *BridgeSelection           `json:"bridge,omitempty"`
	ClientReference  *string                    `json:"client_reference,omitempty"`
	Destination      Destination                `json:"destination"`
	FeeTreatment     string                     `json:"fee_treatment"`
	SettlementMode   string                     `json:"settlement_mode"`
	SourceAccountRef Reference                  `json:"source_account_ref"`
	SourceAsset      Asset                      `json:"source_asset"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *QuoteRequest) UnmarshalJSON(raw []byte) error {
	type plain QuoteRequest
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "amount")
	delete(extra, "bridge")
	delete(extra, "client_reference")
	delete(extra, "destination")
	delete(extra, "fee_treatment")
	delete(extra, "settlement_mode")
	delete(extra, "source_account_ref")
	delete(extra, "source_asset")
	p.AdditionalFields = extra
	*v = QuoteRequest(p)
	return nil
}
func (v QuoteRequest) MarshalJSON() ([]byte, error) {
	type plain QuoteRequest
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "amount":
			continue
		case "bridge":
			continue
		case "client_reference":
			continue
		case "destination":
			continue
		case "fee_treatment":
			continue
		case "settlement_mode":
			continue
		case "source_account_ref":
			continue
		case "source_asset":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type ReceivingAddress struct {
	AccountRef          Reference                  `json:"account_ref"`
	Address             string                     `json:"address"`
	Asset               Asset                      `json:"asset"`
	CreatedAt           *Timestamp                 `json:"created_at,omitempty"`
	Memo                *string                    `json:"memo,omitempty"`
	ReceivingAddressRef Reference                  `json:"receiving_address_ref"`
	Status              string                     `json:"status"`
	Version             Version                    `json:"version"`
	AdditionalFields    map[string]json.RawMessage `json:"-"`
}

func (v *ReceivingAddress) UnmarshalJSON(raw []byte) error {
	type plain ReceivingAddress
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "account_ref")
	delete(extra, "address")
	delete(extra, "asset")
	delete(extra, "created_at")
	delete(extra, "memo")
	delete(extra, "receiving_address_ref")
	delete(extra, "status")
	delete(extra, "version")
	p.AdditionalFields = extra
	*v = ReceivingAddress(p)
	return nil
}
func (v ReceivingAddress) MarshalJSON() ([]byte, error) {
	type plain ReceivingAddress
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "account_ref":
			continue
		case "address":
			continue
		case "asset":
			continue
		case "created_at":
			continue
		case "memo":
			continue
		case "receiving_address_ref":
			continue
		case "status":
			continue
		case "version":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type ReceivingAddressList struct {
	Items            []ReceivingAddress         `json:"items"`
	NextCursor       *string                    `json:"next_cursor,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *ReceivingAddressList) UnmarshalJSON(raw []byte) error {
	type plain ReceivingAddressList
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "items")
	delete(extra, "next_cursor")
	p.AdditionalFields = extra
	*v = ReceivingAddressList(p)
	return nil
}
func (v ReceivingAddressList) MarshalJSON() ([]byte, error) {
	type plain ReceivingAddressList
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "items":
			continue
		case "next_cursor":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type ReceivingAttribution struct {
	Address             string                     `json:"address"`
	CapturedAt          Timestamp                  `json:"captured_at"`
	Memo                *string                    `json:"memo,omitempty"`
	ReceivingAddressRef *Reference                 `json:"receiving_address_ref,omitempty"`
	RouteVersion        *Version                   `json:"route_version,omitempty"`
	AdditionalFields    map[string]json.RawMessage `json:"-"`
}

func (v *ReceivingAttribution) UnmarshalJSON(raw []byte) error {
	type plain ReceivingAttribution
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "address")
	delete(extra, "captured_at")
	delete(extra, "memo")
	delete(extra, "receiving_address_ref")
	delete(extra, "route_version")
	p.AdditionalFields = extra
	*v = ReceivingAttribution(p)
	return nil
}
func (v ReceivingAttribution) MarshalJSON() ([]byte, error) {
	type plain ReceivingAttribution
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "address":
			continue
		case "captured_at":
			continue
		case "memo":
			continue
		case "receiving_address_ref":
			continue
		case "route_version":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type Reference = string
type StatementEntry struct {
	AccountAmount    Money                      `json:"account_amount"`
	Asset            Asset                      `json:"asset"`
	Description      *string                    `json:"description,omitempty"`
	EffectRefs       []Reference                `json:"effect_refs"`
	EntryRef         Reference                  `json:"entry_ref"`
	EntryType        string                     `json:"entry_type"`
	EventTime        Timestamp                  `json:"event_time"`
	Lines            []PostingLine              `json:"lines"`
	Version          Version                    `json:"version"`
	Visibility       *string                    `json:"visibility,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *StatementEntry) UnmarshalJSON(raw []byte) error {
	type plain StatementEntry
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "account_amount")
	delete(extra, "asset")
	delete(extra, "description")
	delete(extra, "effect_refs")
	delete(extra, "entry_ref")
	delete(extra, "entry_type")
	delete(extra, "event_time")
	delete(extra, "lines")
	delete(extra, "version")
	delete(extra, "visibility")
	p.AdditionalFields = extra
	*v = StatementEntry(p)
	return nil
}
func (v StatementEntry) MarshalJSON() ([]byte, error) {
	type plain StatementEntry
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "account_amount":
			continue
		case "asset":
			continue
		case "description":
			continue
		case "effect_refs":
			continue
		case "entry_ref":
			continue
		case "entry_type":
			continue
		case "event_time":
			continue
		case "lines":
			continue
		case "version":
			continue
		case "visibility":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type StatementList struct {
	Items            []StatementSummary         `json:"items"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *StatementList) UnmarshalJSON(raw []byte) error {
	type plain StatementList
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "items")
	p.AdditionalFields = extra
	*v = StatementList(p)
	return nil
}
func (v StatementList) MarshalJSON() ([]byte, error) {
	type plain StatementList
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "items":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type StatementSummary struct {
	AccountAmount    Money                      `json:"account_amount"`
	Asset            Asset                      `json:"asset"`
	Description      *string                    `json:"description,omitempty"`
	EffectRefs       []Reference                `json:"effect_refs"`
	EntryRef         Reference                  `json:"entry_ref"`
	EntryType        string                     `json:"entry_type"`
	EventTime        Timestamp                  `json:"event_time"`
	Version          Version                    `json:"version"`
	Visibility       *string                    `json:"visibility,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *StatementSummary) UnmarshalJSON(raw []byte) error {
	type plain StatementSummary
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "account_amount")
	delete(extra, "asset")
	delete(extra, "description")
	delete(extra, "effect_refs")
	delete(extra, "entry_ref")
	delete(extra, "entry_type")
	delete(extra, "event_time")
	delete(extra, "version")
	delete(extra, "visibility")
	p.AdditionalFields = extra
	*v = StatementSummary(p)
	return nil
}
func (v StatementSummary) MarshalJSON() ([]byte, error) {
	type plain StatementSummary
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "account_amount":
			continue
		case "asset":
			continue
		case "description":
			continue
		case "effect_refs":
			continue
		case "entry_ref":
			continue
		case "entry_type":
			continue
		case "event_time":
			continue
		case "version":
			continue
		case "visibility":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type SubmitTransfer struct {
	ClientReference  *string                    `json:"client_reference,omitempty"`
	QuoteRef         Reference                  `json:"quote_ref"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *SubmitTransfer) UnmarshalJSON(raw []byte) error {
	type plain SubmitTransfer
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "client_reference")
	delete(extra, "quote_ref")
	p.AdditionalFields = extra
	*v = SubmitTransfer(p)
	return nil
}
func (v SubmitTransfer) MarshalJSON() ([]byte, error) {
	type plain SubmitTransfer
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "client_reference":
			continue
		case "quote_ref":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type Timestamp = string
type Transfer struct {
	Attention        *Attention                 `json:"attention,omitempty"`
	Bridges          *[]BridgeStatus            `json:"bridges,omitempty"`
	ClientReference  *string                    `json:"client_reference,omitempty"`
	CompletedAt      *Timestamp                 `json:"completed_at,omitempty"`
	CreatedAt        Timestamp                  `json:"created_at"`
	Executions       *[]ExecutionObservation    `json:"executions,omitempty"`
	Quote            Quote                      `json:"quote"`
	QuoteRef         Reference                  `json:"quote_ref"`
	State            string                     `json:"state"`
	TransferRef      Reference                  `json:"transfer_ref"`
	UpdatedAt        Timestamp                  `json:"updated_at"`
	Version          Version                    `json:"version"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *Transfer) UnmarshalJSON(raw []byte) error {
	type plain Transfer
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "attention")
	delete(extra, "bridges")
	delete(extra, "client_reference")
	delete(extra, "completed_at")
	delete(extra, "created_at")
	delete(extra, "executions")
	delete(extra, "quote")
	delete(extra, "quote_ref")
	delete(extra, "state")
	delete(extra, "transfer_ref")
	delete(extra, "updated_at")
	delete(extra, "version")
	p.AdditionalFields = extra
	*v = Transfer(p)
	return nil
}
func (v Transfer) MarshalJSON() ([]byte, error) {
	type plain Transfer
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "attention":
			continue
		case "bridges":
			continue
		case "client_reference":
			continue
		case "completed_at":
			continue
		case "created_at":
			continue
		case "executions":
			continue
		case "quote":
			continue
		case "quote_ref":
			continue
		case "state":
			continue
		case "transfer_ref":
			continue
		case "updated_at":
			continue
		case "version":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type TransferEvent struct {
	Data             Transfer                   `json:"data"`
	DeliveryId       string                     `json:"delivery_id"`
	EventId          string                     `json:"event_id"`
	OccurredAt       Timestamp                  `json:"occurred_at"`
	Type             string                     `json:"type"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *TransferEvent) UnmarshalJSON(raw []byte) error {
	type plain TransferEvent
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "data")
	delete(extra, "delivery_id")
	delete(extra, "event_id")
	delete(extra, "occurred_at")
	delete(extra, "type")
	p.AdditionalFields = extra
	*v = TransferEvent(p)
	return nil
}
func (v TransferEvent) MarshalJSON() ([]byte, error) {
	type plain TransferEvent
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "data":
			continue
		case "delivery_id":
			continue
		case "event_id":
			continue
		case "occurred_at":
			continue
		case "type":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type TransferInstructions struct {
	Bridge           *BridgeSelection           `json:"bridge,omitempty"`
	Destination      Destination                `json:"destination"`
	SourceAccountRef Reference                  `json:"source_account_ref"`
	SourceAsset      Asset                      `json:"source_asset"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *TransferInstructions) UnmarshalJSON(raw []byte) error {
	type plain TransferInstructions
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "bridge")
	delete(extra, "destination")
	delete(extra, "source_account_ref")
	delete(extra, "source_asset")
	p.AdditionalFields = extra
	*v = TransferInstructions(p)
	return nil
}
func (v TransferInstructions) MarshalJSON() ([]byte, error) {
	type plain TransferInstructions
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "bridge":
			continue
		case "destination":
			continue
		case "source_account_ref":
			continue
		case "source_asset":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type TransferList struct {
	Items            []Transfer                 `json:"items"`
	NextCursor       *string                    `json:"next_cursor,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *TransferList) UnmarshalJSON(raw []byte) error {
	type plain TransferList
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "items")
	delete(extra, "next_cursor")
	p.AdditionalFields = extra
	*v = TransferList(p)
	return nil
}
func (v TransferList) MarshalJSON() ([]byte, error) {
	type plain TransferList
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "items":
			continue
		case "next_cursor":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type TransferOrigin struct {
	ClientReference  *string                    `json:"client_reference,omitempty"`
	TransferRef      Reference                  `json:"transfer_ref"`
	Type             string                     `json:"type"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
}

func (v *TransferOrigin) UnmarshalJSON(raw []byte) error {
	type plain TransferOrigin
	var p plain
	if err := json.Unmarshal(raw, &p); err != nil {
		return err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(raw, &extra); err != nil {
		return err
	}
	delete(extra, "client_reference")
	delete(extra, "transfer_ref")
	delete(extra, "type")
	p.AdditionalFields = extra
	*v = TransferOrigin(p)
	return nil
}
func (v TransferOrigin) MarshalJSON() ([]byte, error) {
	type plain TransferOrigin
	raw, err := json.Marshal(plain(v))
	if err != nil {
		return nil, err
	}
	var out map[string]json.RawMessage
	if err = json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	for k, value := range v.AdditionalFields {
		switch k {
		case "client_reference":
			continue
		case "transfer_ref":
			continue
		case "type":
			continue
		}
		if _, exists := out[k]; !exists {
			out[k] = value
		}
	}
	return json.Marshal(out)
}

type Version = string

type Client struct {
	BaseURL string
	HTTP    *http.Client
	Bearer  string
}
type ResponseError = transport.ResponseError

var ErrResponseTooLarge = transport.ErrResponseTooLarge

func (c *Client) call(ctx context.Context, method, path string, query url.Values, key string, body any, out any) error {
	var raw json.RawMessage
	if body != nil {
		var err error
		raw, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}
	client := http.Client{}
	if c.HTTP != nil {
		client = *c.HTTP
	}
	boundary, err := transport.New(c.BaseURL, client)
	if err != nil {
		return err
	}
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	response, err := boundary.Call(ctx, c.Bearer, method, path, key, raw)
	if err != nil {
		return err
	}
	return json.Unmarshal(response, out)
}
func (c *Client) ListAccounts(ctx context.Context, query url.Values) (*AccountList, error) {
	path := "/v1/accounts"
	out := new(AccountList)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) GetBalance(ctx context.Context, AccountRef string, query url.Values) (*Balance, error) {
	path := "/v1/accounts/{account_ref}/balance"
	path = strings.ReplaceAll(path, "{account_ref}", url.PathEscape(AccountRef))
	out := new(Balance)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) ListStatementEntries(ctx context.Context, AccountRef string, query url.Values) (*StatementList, error) {
	path := "/v1/accounts/{account_ref}/statement-entries"
	path = strings.ReplaceAll(path, "{account_ref}", url.PathEscape(AccountRef))
	out := new(StatementList)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) GetStatementEntry(ctx context.Context, AccountRef string, EntryRef string, query url.Values) (*StatementEntry, error) {
	path := "/v1/accounts/{account_ref}/statement-entries/{entry_ref}"
	path = strings.ReplaceAll(path, "{account_ref}", url.PathEscape(AccountRef))
	path = strings.ReplaceAll(path, "{entry_ref}", url.PathEscape(EntryRef))
	out := new(StatementEntry)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) ListAssets(ctx context.Context, query url.Values) (*AssetList, error) {
	path := "/v1/assets"
	out := new(AssetList)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) ListBridgeRoutes(ctx context.Context, query url.Values) (*BridgeRouteList, error) {
	path := "/v1/bridge-routes"
	out := new(BridgeRouteList)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) ListDepositRevisions(ctx context.Context, query url.Values) (*DepositList, error) {
	path := "/v1/deposits"
	out := new(DepositList)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) GetDeposit(ctx context.Context, DepositRef string, query url.Values) (*Deposit, error) {
	path := "/v1/deposits/{deposit_ref}"
	path = strings.ReplaceAll(path, "{deposit_ref}", url.PathEscape(DepositRef))
	out := new(Deposit)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) ListFinancialEffects(ctx context.Context, query url.Values) (*FinancialEffectList, error) {
	path := "/v1/financial-effects"
	out := new(FinancialEffectList)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) GetFinancialEffect(ctx context.Context, EffectRef string, query url.Values) (*FinancialEffect, error) {
	path := "/v1/financial-effects/{effect_ref}"
	path = strings.ReplaceAll(path, "{effect_ref}", url.PathEscape(EffectRef))
	out := new(FinancialEffect)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) CreateQuote(ctx context.Context, query url.Values, key string, body QuoteRequest) (*Quote, error) {
	path := "/v1/quotes"
	out := new(Quote)
	err := c.call(ctx, "POST", path, query, key, body, out)
	return out, err
}
func (c *Client) ListReceivingAddresses(ctx context.Context, query url.Values) (*ReceivingAddressList, error) {
	path := "/v1/receiving-addresses"
	out := new(ReceivingAddressList)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) CreateReceivingAddress(ctx context.Context, query url.Values, key string, body CreateReceivingAddress) (*ReceivingAddress, error) {
	path := "/v1/receiving-addresses"
	out := new(ReceivingAddress)
	err := c.call(ctx, "POST", path, query, key, body, out)
	return out, err
}
func (c *Client) GetReceivingAddress(ctx context.Context, ReceivingAddressRef string, query url.Values) (*ReceivingAddress, error) {
	path := "/v1/receiving-addresses/{receiving_address_ref}"
	path = strings.ReplaceAll(path, "{receiving_address_ref}", url.PathEscape(ReceivingAddressRef))
	out := new(ReceivingAddress)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) ListTransfers(ctx context.Context, query url.Values) (*TransferList, error) {
	path := "/v1/transfers"
	out := new(TransferList)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
func (c *Client) SubmitTransfer(ctx context.Context, query url.Values, key string, body SubmitTransfer) (*Transfer, error) {
	path := "/v1/transfers"
	out := new(Transfer)
	err := c.call(ctx, "POST", path, query, key, body, out)
	return out, err
}
func (c *Client) GetTransfer(ctx context.Context, TransferRef string, query url.Values) (*Transfer, error) {
	path := "/v1/transfers/{transfer_ref}"
	path = strings.ReplaceAll(path, "{transfer_ref}", url.PathEscape(TransferRef))
	out := new(Transfer)
	err := c.call(ctx, "GET", path, query, "", nil, out)
	return out, err
}
