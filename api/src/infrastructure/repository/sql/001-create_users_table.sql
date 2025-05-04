CREATE TABLE public.users (
	"id" uuid PRIMARY KEY DEFAULT public.uuid_generate_v4() NOT NULL,
    "email" text NULL,
	"name" text NOT NULL,
	"password" text NOT NULL,
    "created_at" timestamptz DEFAULT now(),
    "updated_at" timestamptz NULL,
    "deleted_at" timestamptz NULL
)
;

CREATE UNIQUE INDEX index_users_email ON users (email) WHERE (deleted_at IS NULL)
;

CREATE TRIGGER trigger_users_updated_at
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION public.update_updated_at()
;
