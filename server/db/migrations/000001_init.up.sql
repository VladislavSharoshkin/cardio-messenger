-- Table: public.syncs
CREATE TABLE IF NOT EXISTS public.syncs
(
    id bigserial,
    query text NOT NULL,
    type bigint NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    CONSTRAINT syncs_pkey PRIMARY KEY (id)
);

-- Table: public.users
CREATE TABLE IF NOT EXISTS public.users
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    login text NOT NULL,
    hash_pass text,
    first_name text NOT NULL,
    middle_name text,
    last_name text NOT NULL,
    avatar_id bigint,
    about text,
    sync_id bigint,
    sync_remote_id text,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT fk_users_sync FOREIGN KEY (sync_id)
        REFERENCES public.syncs (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS users_deleted_at_idx
    ON public.users USING btree
    (deleted_at ASC NULLS LAST, login ASC NULLS LAST);

-- Table: public.files
CREATE TABLE IF NOT EXISTS public.files
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    name text NOT NULL,
    size bigint NOT NULL,
    hash text NOT NULL,
    type bigint NOT NULL,
    creator_id bigint NOT NULL,
    token text NOT NULL,
    CONSTRAINT files_pkey PRIMARY KEY (id)
);

ALTER TABLE users
    ADD CONSTRAINT fk_users_avatar FOREIGN KEY (avatar_id)
        REFERENCES public.files (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE;

ALTER TABLE files
    ADD CONSTRAINT fk_files_creator FOREIGN KEY (creator_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE;

-- Table: public.branches
CREATE TABLE IF NOT EXISTS public.branches
(
    id bigserial,
    name text NOT NULL,
    sync_id bigint,
    sync_remote_id text,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    CONSTRAINT branches_pkey PRIMARY KEY (id),
    CONSTRAINT branches_name_key UNIQUE (name)
);

-- Index: idx_Branch_1
CREATE UNIQUE INDEX IF NOT EXISTS "idx_Branch_1"
    ON public.branches USING btree
        (sync_id ASC NULLS LAST, sync_remote_id ASC NULLS LAST);

-- Table: public.staffs
CREATE TABLE IF NOT EXISTS public.staffs
(
    id bigserial,
    name text NOT NULL,
    sync_id bigint,
    sync_remote_id text,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    CONSTRAINT staffs_pkey PRIMARY KEY (id),
    CONSTRAINT staffs_name_key UNIQUE (name)
);

CREATE UNIQUE INDEX IF NOT EXISTS "staffs_deleted_at_idx"
    ON public.staffs USING btree
        (sync_id ASC NULLS LAST, sync_remote_id ASC NULLS LAST);

-- Table: public.user_staffs
CREATE TABLE IF NOT EXISTS public.user_staffs
(
    id bigserial,
    user_id bigint NOT NULL,
    branch_id bigint NOT NULL,
    staff_id bigint NOT NULL,
    sync_id bigint,
    sync_remote_id text,
    sync_remote_user_id text,
    sync_remote_branch_id text,
    sync_remote_staff_id text,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    CONSTRAINT user_staffs_pkey PRIMARY KEY (id),
    CONSTRAINT fk_user_staffs_branch FOREIGN KEY (branch_id)
        REFERENCES public.branches (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT fk_user_staffs_staff FOREIGN KEY (staff_id)
        REFERENCES public.staffs (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT fk_users_user_staff_list FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "user_staffs_deleted_at_idx"
    ON public.user_staffs USING btree
    (deleted_at ASC NULLS LAST, user_id ASC NULLS LAST, branch_id ASC NULLS LAST, staff_id ASC NULLS LAST);

-- Table: public.tokens
CREATE TABLE IF NOT EXISTS public.tokens
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    user_id bigint NOT NULL,
    token text NOT NULL,
    push text,
    CONSTRAINT tokens_pkey PRIMARY KEY (id),
    CONSTRAINT tokens_token_key UNIQUE (token),
    CONSTRAINT fk_users_token_list FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);



-- Table: public.subscribers
CREATE TABLE IF NOT EXISTS public.subscribers
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    user_id bigint NOT NULL,
    subscriber_id bigint NOT NULL,
    CONSTRAINT subscribers_pk PRIMARY KEY (id),
    CONSTRAINT subscribers_fk FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT subscribers_fk_1 FOREIGN KEY (subscriber_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS subscribers_deleted_at_idx
    ON public.subscribers USING btree
    (deleted_at ASC NULLS LAST, user_id ASC NULLS LAST, subscriber_id ASC NULLS LAST);



CREATE TABLE IF NOT EXISTS public.server_pushs
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    token text NOT NULL,
    push text NOT NULL,
    CONSTRAINT tokens_pkey_8 PRIMARY KEY (id),
    CONSTRAINT tokens_token_key_5 UNIQUE (token)
);


-- Table: public.chats
CREATE TABLE IF NOT EXISTS public.chats
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    creator_id bigint NOT NULL,
    type bigint NOT NULL,
    name text,
    type_unique bigint,
    avatar_id bigint,
    about text,
    name_unique text,
    CONSTRAINT chats_pkey PRIMARY KEY (id),
    CONSTRAINT chats_type_unique_key UNIQUE (type_unique),
    CONSTRAINT fk_chats_avatar FOREIGN KEY (avatar_id)
        REFERENCES public.files (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT fk_chats_creator FOREIGN KEY (creator_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT chats_check_type CHECK (type >= 1 AND type <= 2)
);

-- Table: public.messages
CREATE TABLE IF NOT EXISTS public.messages
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    sender_id bigint NOT NULL,
    chat_id bigint NOT NULL,
    text text,
    CONSTRAINT messages_pkey PRIMARY KEY (id),
    CONSTRAINT fk_chats_last_message FOREIGN KEY (chat_id)
        REFERENCES public.chats (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT fk_messages_sender FOREIGN KEY (sender_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);


-- Table: public.reads
CREATE TABLE IF NOT EXISTS public.reads
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    message_id bigint NOT NULL,
    user_id bigint NOT NULL,
    CONSTRAINT message_files_pkey_2 PRIMARY KEY (id),
    CONSTRAINT reads_fk FOREIGN KEY (message_id)
        REFERENCES public.messages (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT reads_fk_1 FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS reads_deleted_at_idx
    ON public.reads USING btree
    (deleted_at ASC NULLS LAST, message_id ASC NULLS LAST, user_id ASC NULLS LAST);

-- Table: public.posts
CREATE TABLE IF NOT EXISTS public.posts
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    user_id bigint NOT NULL,
    CONSTRAINT tokens_pkey_3 PRIMARY KEY (id),
    CONSTRAINT posts_fk FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);

-- Table: public.participants
CREATE TABLE IF NOT EXISTS public.participants
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    chat_id bigint NOT NULL,
    user_id bigint NOT NULL,
    last_read_message_id bigint,
    CONSTRAINT participants_pkey PRIMARY KEY (id),
    CONSTRAINT fk_chats_participant_list FOREIGN KEY (chat_id)
        REFERENCES public.chats (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT fk_participants_last_read_message FOREIGN KEY (last_read_message_id)
        REFERENCES public.messages (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT fk_participants_user FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS participants_deleted_at_idx
    ON public.participants USING btree
    (deleted_at ASC NULLS LAST, chat_id ASC NULLS LAST, user_id ASC NULLS LAST);

-- Table: public.onlines
CREATE TABLE IF NOT EXISTS public.onlines
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    user_id bigint NOT NULL,
    type bigint NOT NULL,
    CONSTRAINT tokens_pkey_1_1 PRIMARY KEY (id),
    CONSTRAINT onlines_fk FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT onlines_check_type CHECK (type >= 1 AND type <= 2)
);


-- Table: public.message_files
CREATE TABLE IF NOT EXISTS public.message_files
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    message_id bigint NOT NULL,
    file_id bigint NOT NULL,
    CONSTRAINT message_files_pkey PRIMARY KEY (id),
    CONSTRAINT fk_message_files_file FOREIGN KEY (file_id)
        REFERENCES public.files (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT fk_messages_message_files FOREIGN KEY (message_id)
        REFERENCES public.messages (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);

-- Table: public.logs
CREATE TABLE IF NOT EXISTS public.logs
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    type bigint NOT NULL,
    text text NOT NULL,
    remote_addr text,
    request_uri text,
    error_key bigint,
    CONSTRAINT tokens_pkey_4 PRIMARY KEY (id)
);

-- Table: public.likes
CREATE TABLE IF NOT EXISTS public.likes
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    user_id bigint NOT NULL,
    post_id bigint,
    CONSTRAINT tokens_pkey_2 PRIMARY KEY (id),
    CONSTRAINT likes_fk FOREIGN KEY (post_id)
        REFERENCES public.posts (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT likes_fk2 FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);

-- Table: public.forward_messages
CREATE TABLE IF NOT EXISTS public.forward_messages
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    message_id bigint NOT NULL,
    forward_message_id bigint NOT NULL,
    CONSTRAINT message_files_pkey_1 PRIMARY KEY (id),
    CONSTRAINT forward_messages_fk FOREIGN KEY (message_id)
        REFERENCES public.messages (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT forward_messages_fk_1 FOREIGN KEY (forward_message_id)
        REFERENCES public.messages (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);


-- Table: public.calls



CREATE TABLE IF NOT EXISTS public.calls
(
    id bigserial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    sync_type bigint,
    sync_remote_id text,
    socket_id text,
    CONSTRAINT calls_pkey PRIMARY KEY (id)
);




-- Table: public.activitys

-- DROP TABLE IF EXISTS public.activitys;

CREATE TABLE IF NOT EXISTS public.activitys
(
    id bigserial,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone NOT NULL DEFAULT '9999-01-01 00:00:00+04'::timestamp with time zone,
    user_id bigint NOT NULL,
    chat_id bigint NOT NULL,
    type bigint NOT NULL,
    CONSTRAINT tokens_pkey_1 PRIMARY KEY (id),
    CONSTRAINT activitys_fk FOREIGN KEY (chat_id)
        REFERENCES public.chats (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT activitys_fk_1 FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE,
    CONSTRAINT activitys_check_type CHECK (type >= 1 AND type <= 2)
);
