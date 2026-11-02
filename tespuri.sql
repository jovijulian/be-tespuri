/*
 Navicat Premium Dump SQL

 Source Server Type    : PostgreSQL
 Source Catalog        : tespuri
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 150019 (150019)
 File Encoding         : 65001

 Date: 02/11/2026 23:15:40
*/


-- ----------------------------
-- Sequence structure for entities_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."entities_id_seq";
CREATE SEQUENCE "public"."entities_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."entities_id_seq" OWNER TO "reeval_admin";

-- ----------------------------
-- Table structure for entities
-- ----------------------------
DROP TABLE IF EXISTS "public"."entities";
CREATE TABLE "public"."entities" (
  "id" int8 NOT NULL DEFAULT nextval('entities_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "name" varchar(150) COLLATE "pg_catalog"."default" NOT NULL,
  "code" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "category" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
  "status" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'active'::character varying,
  "latitude" numeric(10,8) NOT NULL,
  "longitude" numeric(11,8) NOT NULL,
  "description" text COLLATE "pg_catalog"."default",
  "metadata" jsonb DEFAULT '{}'::jsonb
)
;
ALTER TABLE "public"."entities" OWNER TO "reeval_admin";

-- ----------------------------
-- Records of entities
-- ----------------------------
BEGIN;
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (2, '2026-11-02 22:51:42.804471+07', '2026-11-02 22:51:42.804471+07', NULL, 'Intercity Heavy Truck 04', 'VEH-JKT-404', 'vehicle', 'active', -6.20876300, 106.84559900, 'Freight hauler currently on Jakarta - Cikampek logistics corridor', '{"speed_kmh": 62, "payload_ton": 12.4, "plate_number": "B 9876 XYZ"}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (3, '2026-11-02 22:51:43.121448+07', '2026-11-02 22:51:43.121448+07', NULL, 'Field Support Patrol 02', 'VEH-BDG-202', 'vehicle', 'maintenance', -6.89148000, 107.61065900, 'Field maintenance patrol vehicle undergoing engine tune-up', '{"workshop": "Dago Central Garage", "plate_number": "D 4512 EF", "estimated_finish": "2026-11-03"}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (4, '2026-11-02 22:51:43.371722+07', '2026-11-02 22:51:43.371722+07', NULL, 'Air Quality Station Dago Atas', 'IOT-AQ-101', 'iot_device', 'active', -6.87948000, 107.61520000, 'Telemetry node monitoring ambient air index, PM2.5, and humidity', '{"pm25": 22.4, "firmware": "v2.1.0", "sensor_type": "Laser Particle AQI", "temperature": 23.8}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (5, '2026-11-02 22:51:43.688497+07', '2026-11-02 22:51:43.688497+07', NULL, 'Weather Station Gedung Sate', 'IOT-WS-201', 'iot_device', 'active', -6.90248100, 107.61876500, 'High-precision microclimate telemetry monitoring wind and precipitation', '{"humidity": 78, "wind_speed": 12.1, "rainfall_mm": 0}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (6, '2026-11-02 22:51:44.439389+07', '2026-11-02 22:51:44.439389+07', NULL, 'Water Level Sensor Citarum Dayeuhkolot', 'IOT-WL-301', 'iot_device', 'active', -6.98561200, 107.62534100, 'River surge sensor for automated flood early warning', '{"telemetry": "LoRaWAN", "alert_level": "NORMAL", "water_level_m": 3.42}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (7, '2026-11-02 22:51:45.054506+07', '2026-11-02 22:51:45.054506+07', NULL, 'Smart Solar Inverter Sensor Bandung', 'IOT-PV-401', 'iot_device', 'maintenance', -6.93441200, 107.60412300, 'Rooftop solar generator telemetry node with inverter communication', '{"power_kw": 0, "fault_code": "E-INV-04", "technician": "Agus Santoso"}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (8, '2026-11-02 22:51:46.142288+07', '2026-11-02 22:51:46.142288+07', NULL, 'IoT Gateway Sudirman Hub', 'IOT-GW-501', 'iot_device', 'active', -6.21550000, 106.82010000, 'Edge aggregation gateway for regional smart street lighting nodes', '{"uptime_days": 124, "connected_nodes": 48, "packet_loss_pct": 0.02}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (9, '2026-11-02 22:51:47.151842+07', '2026-11-02 22:51:47.151842+07', NULL, 'PT LEN Industri HQ Facility', 'FAC-LEN-001', 'facility', 'active', -6.94314200, 107.60832100, 'Corporate headquarters and R&D manufacturing engineering campus', '{"type": "Headquarters & Lab", "address": "Jl. Soekarno Hatta No.442, Bandung", "capacity": 1200}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (10, '2026-11-02 22:51:47.676605+07', '2026-11-02 22:51:47.676605+07', NULL, 'Logistics Distribution Center Soekarno-Hatta', 'FAC-DC-002', 'facility', 'active', -6.94821000, 107.67123000, 'Central warehouse and regional sorting fulfillment facility', '{"dock_doors": 14, "storage_m2": 8500, "cold_storage": true}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (11, '2026-11-02 22:51:47.936074+07', '2026-11-02 22:51:47.936074+07', NULL, 'Greater Jakarta Regional Hub', 'FAC-HUB-003', 'facility', 'active', -6.22550000, 106.80050000, 'Regional commercial and service integration operations hub', '{"status": "Operational", "metro_coverage": "Jabodetabek"}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (12, '2026-11-02 22:51:48.481335+07', '2026-11-02 22:51:48.481335+07', NULL, 'Depot & Workshop Kiaracondong', 'FAC-DEP-004', 'facility', 'maintenance', -6.92415000, 107.64789000, 'Fleet maintenance workshop and spare parts depot', '{"active_tickets": 4, "maintenance_bays": 6}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (13, '2026-11-02 22:52:45.135469+07', '2026-11-02 22:52:49.970997+07', '2026-11-02 22:52:55.025364+07', 'Autonomous Drone Station Alpha', 'IOT-DRONE-01', 'iot_device', 'maintenance', -6.91250000, 107.61500000, 'Undergoing routine motor replacement', '{"docked": true, "battery_level": 90}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (14, '2026-11-02 23:10:33.654036+07', '2026-11-02 23:10:33.654036+07', NULL, 'Smart Weather Telemetry Hub 09', 'IOT-WS-909', 'iot_device', 'active', -6.91474400, 107.60981000, 'Urban microclimate and air sensor on municipal light pole', '{"firmware": "v2.4.1", "battery_pct": 98, "humidity_pct": 72, "temperature_c": 24.5}');
INSERT INTO "public"."entities" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "category", "status", "latitude", "longitude", "description", "metadata") VALUES (1, '2026-11-02 22:51:42.55592+07', '2026-11-02 23:12:12.736565+07', '2026-11-02 23:12:42.629584+07', 'Logistics Fleet Van 01 (Updated)', 'VEH-BDG-101', 'vehicle', 'maintenance', -6.92000000, 107.62500000, 'Vehicle undergoing tire replacement', '{"workshop": "Central Depot", "plate_number": "D 1024 AA"}');
COMMIT;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."entities_id_seq"
OWNED BY "public"."entities"."id";
SELECT setval('"public"."entities_id_seq"', 14, true);

-- ----------------------------
-- Indexes structure for table entities
-- ----------------------------
CREATE INDEX "idx_entities_category" ON "public"."entities" USING btree (
  "category" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "idx_entities_code" ON "public"."entities" USING btree (
  "code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_entities_coords" ON "public"."entities" USING btree (
  "latitude" "pg_catalog"."numeric_ops" ASC NULLS LAST,
  "longitude" "pg_catalog"."numeric_ops" ASC NULLS LAST
);
CREATE INDEX "idx_entities_deleted_at" ON "public"."entities" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_entities_latitude" ON "public"."entities" USING btree (
  "latitude" "pg_catalog"."numeric_ops" ASC NULLS LAST
);
CREATE INDEX "idx_entities_longitude" ON "public"."entities" USING btree (
  "longitude" "pg_catalog"."numeric_ops" ASC NULLS LAST
);
CREATE INDEX "idx_entities_status" ON "public"."entities" USING btree (
  "status" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Checks structure for table entities
-- ----------------------------
ALTER TABLE "public"."entities" ADD CONSTRAINT "chk_entities_latitude" CHECK (latitude >= '-90.0'::numeric AND latitude <= 90.0);
ALTER TABLE "public"."entities" ADD CONSTRAINT "chk_entities_longitude" CHECK (longitude >= '-180.0'::numeric AND longitude <= 180.0);
ALTER TABLE "public"."entities" ADD CONSTRAINT "chk_entities_category" CHECK (category::text = ANY (ARRAY['vehicle'::character varying, 'iot_device'::character varying, 'facility'::character varying]::text[]));
ALTER TABLE "public"."entities" ADD CONSTRAINT "chk_entities_status" CHECK (status::text = ANY (ARRAY['active'::character varying, 'inactive'::character varying, 'maintenance'::character varying]::text[]));

-- ----------------------------
-- Primary Key structure for table entities
-- ----------------------------
ALTER TABLE "public"."entities" ADD CONSTRAINT "entities_pkey" PRIMARY KEY ("id");
