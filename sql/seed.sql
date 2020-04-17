--
-- PostgreSQL database dump
--

-- Dumped from database version 10.2
-- Dumped by pg_dump version 10.2

-- Started on 2020-04-17 12:39:18 GMT

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public, pg_catalog;

--
-- TOC entry 3338 (class 0 OID 62932)
-- Dependencies: 200
-- Data for Name: address_type; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3339 (class 0 OID 62942)
-- Dependencies: 201
-- Data for Name: addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3344 (class 0 OID 63016)
-- Dependencies: 206
-- Data for Name: cart_type; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3345 (class 0 OID 63025)
-- Dependencies: 207
-- Data for Name: carts; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3340 (class 0 OID 62957)
-- Dependencies: 202
-- Data for Name: customer_address; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3343 (class 0 OID 62997)
-- Dependencies: 205
-- Data for Name: customer_payment_types; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3337 (class 0 OID 62920)
-- Dependencies: 199
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3347 (class 0 OID 63054)
-- Dependencies: 209
-- Data for Name: delivery_status; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3360 (class 0 OID 63286)
-- Dependencies: 222
-- Data for Name: gopg_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3354 (class 0 OID 63159)
-- Dependencies: 216
-- Data for Name: inventory; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3357 (class 0 OID 63183)
-- Dependencies: 219
-- Data for Name: oauth_clients; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO oauth_clients (id, secret, domain, data) VALUES ('4x45RYbu21vYFHABaPjl', 'cbmDBX2LYiwvlsJntdpa', 'http://127.0.0.1:8080', '{}');


--
-- TOC entry 3352 (class 0 OID 63116)
-- Dependencies: 214
-- Data for Name: order_items; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3346 (class 0 OID 63044)
-- Dependencies: 208
-- Data for Name: order_status; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3348 (class 0 OID 63064)
-- Dependencies: 210
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3341 (class 0 OID 62978)
-- Dependencies: 203
-- Data for Name: payment_types; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3342 (class 0 OID 62989)
-- Dependencies: 204
-- Data for Name: payment_vendor; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3353 (class 0 OID 63140)
-- Dependencies: 215
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3350 (class 0 OID 63095)
-- Dependencies: 212
-- Data for Name: product_brands; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3356 (class 0 OID 63172)
-- Dependencies: 218
-- Data for Name: product_categories; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3351 (class 0 OID 63102)
-- Dependencies: 213
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3335 (class 0 OID 62892)
-- Dependencies: 197
-- Data for Name: user_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO user_roles (id, role_name, description, comment, updated_by, created_at, updated_at, deleted_at) VALUES (1, 'admin', 'Administrator', NULL, NULL, '2020-03-02 12:49:31.386119+00', NULL, NULL);


--
-- TOC entry 3336 (class 0 OID 62903)
-- Dependencies: 198
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO users (id, username, password, firstname, middlename, lastname, gender, email_work, phone_work, email_personal, phone_personal, role, status, last_login, created_at, updated_at, deleted_at, updated_by) VALUES ('63f284c3-1891-4905-a183-57f621aca134', 'admin', '$2a$14$TH23lPu7kA9QiRqW8SCNJOg182LKQ7okjhCThCN.ICSw9dgmBk2a2', 'admin', NULL, 'admin', 'm', 'admin@admin.com', NULL, NULL, NULL, 1, true, NULL, '2020-03-03 13:09:20.037895+00', NULL, NULL, NULL);


--
-- TOC entry 3373 (class 0 OID 0)
-- Dependencies: 211
-- Name: product_brands_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('product_brands_id_seq', 1, false);


--
-- TOC entry 3374 (class 0 OID 0)
-- Dependencies: 217
-- Name: product_categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('product_categories_id_seq', 1, false);


-- Completed on 2020-04-17 12:39:18 GMT

--
-- PostgreSQL database dump complete
--

