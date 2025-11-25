CREATE TABLE IF NOT EXISTS "events"(
   "id" BIGSERIAL PRIMARY KEY,
   "service_id" BIGINT NOT NULL,
   "event_type" SMALLINT NOT NULL,
   "value" JSONB DEFAULT '{}' NOT NULL,
   "status" SMALLINT DEFAULT 0,
   "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   CONSTRAINT fk_events_services_service_id
      FOREIGN KEY(service_id) 
       REFERENCES services(id)
);
