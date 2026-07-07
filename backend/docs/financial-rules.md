# Financial Rules

## Rule 1

Wallet balances must never become negative.

---

## Rule 2

Every successful debit must be recorded as a transaction.

---

## Rule 3

Every purchase creates exactly one order.

---

## Rule 4

Financial operations must be atomic.

Either every operation succeeds or the entire transaction is rolled back.

---

## Rule 5

Only the Wallet Service may modify wallet balances.

Repositories must never change balances directly.

---

## Rule 6

All monetary calculations use decimal.Decimal.

float32 and float64 are prohibited.

---

## Rule 7

Repositories persist data.

Services enforce business rules.

Handlers expose APIs.

Each layer has exactly one responsibility.