-- ============================================================
-- JCK Connect Development Seed Data
-- Products
-- ============================================================

-- ============================================================
-- AIRTIME
-- ============================================================

INSERT INTO products (
    id,
    code,
    name,
    category,
    provider,
    description,
    price,
    currency,
    active,
    created_at,
    updated_at
)
VALUES
(
    '10000000-0000-0000-0000-000000000001',
    'MTN-AIRTIME',
    'MTN Airtime',
    'AIRTIME',
    'VTPASS',
    'MTN Airtime Recharge',
    1.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000002',
    'AIRTEL-AIRTIME',
    'Airtel Airtime',
    'AIRTIME',
    'VTPASS',
    'Airtel Airtime Recharge',
    1.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000003',
    'GLO-AIRTIME',
    'Glo Airtime',
    'AIRTIME',
    'VTPASS',
    'Glo Airtime Recharge',
    1.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000004',
    '9MOBILE-AIRTIME',
    '9mobile Airtime',
    'AIRTIME',
    'VTPASS',
    '9mobile Airtime Recharge',
    1.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
)
ON CONFLICT (code) DO NOTHING;

-- ============================================================
-- DATA
-- ============================================================

INSERT INTO products (
    id,
    code,
    name,
    category,
    provider,
    description,
    price,
    currency,
    active,
    created_at,
    updated_at
)
VALUES
(
    '10000000-0000-0000-0000-000000000005',
    'MTN-DATA-1GB',
    'MTN 1GB Data',
    'DATA',
    'VTPASS',
    'MTN 1GB Data Bundle',
    2.50,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000006',
    'MTN-DATA-2GB',
    'MTN 2GB Data',
    'DATA',
    'VTPASS',
    'MTN 2GB Data Bundle',
    4.50,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000007',
    'AIRTEL-DATA-1GB',
    'Airtel 1GB Data',
    'DATA',
    'VTPASS',
    'Airtel 1GB Data Bundle',
    2.50,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000008',
    'GLO-DATA-1GB',
    'Glo 1GB Data',
    'DATA',
    'VTPASS',
    'Glo 1GB Data Bundle',
    2.50,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000009',
    '9MOBILE-DATA-1GB',
    '9mobile 1GB Data',
    'DATA',
    'VTPASS',
    '9mobile 1GB Data Bundle',
    2.50,
    'PI',
    TRUE,
    NOW(),
    NOW()
)
ON CONFLICT (code) DO NOTHING;

-- ============================================================
-- ELECTRICITY
-- ============================================================

INSERT INTO products (
    id,
    code,
    name,
    category,
    provider,
    description,
    price,
    currency,
    active,
    created_at,
    updated_at
)
VALUES
(
    '10000000-0000-0000-0000-000000000010',
    'EKEDC',
    'Eko Electricity',
    'ELECTRICITY',
    'VTPASS',
    'Eko Electricity Distribution Company',
    5.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000011',
    'IKEDC',
    'Ikeja Electricity',
    'ELECTRICITY',
    'VTPASS',
    'Ikeja Electricity Distribution Company',
    5.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000012',
    'AEDC',
    'Abuja Electricity',
    'ELECTRICITY',
    'VTPASS',
    'Abuja Electricity Distribution Company',
    5.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000013',
    'IBEDC',
    'Ibadan Electricity',
    'ELECTRICITY',
    'VTPASS',
    'Ibadan Electricity Distribution Company',
    5.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000014',
    'PHED',
    'Port Harcourt Electricity',
    'ELECTRICITY',
    'VTPASS',
    'Port Harcourt Electricity Distribution Company',
    5.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
)
ON CONFLICT (code) DO NOTHING;

-- ============================================================
-- TV
-- ============================================================

INSERT INTO products (
    id,
    code,
    name,
    category,
    provider,
    description,
    price,
    currency,
    active,
    created_at,
    updated_at
)
VALUES
(
    '10000000-0000-0000-0000-000000000015',
    'DSTV-COMPACT',
    'DStv Compact',
    'TV',
    'VTPASS',
    'DStv Compact Subscription',
    15.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000016',
    'DSTV-PREMIUM',
    'DStv Premium',
    'TV',
    'VTPASS',
    'DStv Premium Subscription',
    25.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000017',
    'GOTV-MAX',
    'GOtv Max',
    'TV',
    'VTPASS',
    'GOtv Max Subscription',
    10.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
),
(
    '10000000-0000-0000-0000-000000000018',
    'STARTIMES-BASIC',
    'StarTimes Basic',
    'TV',
    'VTPASS',
    'StarTimes Basic Subscription',
    8.00,
    'PI',
    TRUE,
    NOW(),
    NOW()
)
ON CONFLICT (code) DO NOTHING;