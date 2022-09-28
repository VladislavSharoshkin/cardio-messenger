-- Table: public.users

DROP TABLE IF EXISTS public.users CASCADE;

-- Index: idx_users_deleted_at

DROP INDEX IF EXISTS public.idx_users_deleted_at CASCADE;

-- Index: users_deleted_at_idx

DROP INDEX IF EXISTS public.users_deleted_at_idx CASCADE;

-- Table: public.user_staffs

DROP TABLE IF EXISTS public.user_staffs CASCADE;

-- Index: idx_UserStaff_1

DROP INDEX IF EXISTS public."idx_UserStaff_1" CASCADE;

-- Table: public.tokens

DROP TABLE IF EXISTS public.tokens CASCADE;

-- Index: idx_tokens_deleted_at

DROP INDEX IF EXISTS public.idx_tokens_deleted_at CASCADE;

-- Table: public.syncs

DROP TABLE IF EXISTS public.syncs CASCADE;

-- Table: public.subscribers

DROP TABLE IF EXISTS public.subscribers CASCADE;

-- Index: subscribers_deleted_at_idx

DROP INDEX IF EXISTS public.subscribers_deleted_at_idx CASCADE;

-- Table: public.staffs

DROP TABLE IF EXISTS public.staffs CASCADE;

-- Index: idx_Staff_1

DROP INDEX IF EXISTS public."idx_Staff_1" CASCADE;

-- Table: public.server_pushs

DROP TABLE IF EXISTS public.server_pushs CASCADE;

-- Table: public.reads

DROP TABLE IF EXISTS public.reads CASCADE;

-- Index: reads_deleted_at_idx

DROP INDEX IF EXISTS public.reads_deleted_at_idx CASCADE;

-- Table: public.posts

DROP TABLE IF EXISTS public.posts CASCADE;

-- Index: idx_tokens_deleted_at_3

DROP INDEX IF EXISTS public.idx_tokens_deleted_at_3 CASCADE;

-- Table: public.participants

DROP TABLE IF EXISTS public.participants CASCADE;

-- Index: participants_deleted_at_idx

DROP INDEX IF EXISTS public.participants_deleted_at_idx CASCADE;

-- Table: public.onlines

DROP TABLE IF EXISTS public.onlines CASCADE;

-- Index: idx_tokens_deleted_at_1_1

DROP INDEX IF EXISTS public.idx_tokens_deleted_at_1_1 CASCADE;

-- Table: public.messages

DROP TABLE IF EXISTS public.messages CASCADE;

-- Index: idx_messages_deleted_at

DROP INDEX IF EXISTS public.idx_messages_deleted_at CASCADE;

-- Table: public.message_files

DROP TABLE IF EXISTS public.message_files CASCADE;

-- Index: idx_message_files_deleted_at

DROP INDEX IF EXISTS public.idx_message_files_deleted_at CASCADE;

-- Table: public.logs

DROP TABLE IF EXISTS public.logs CASCADE;

-- Index: idx_tokens_deleted_at_4

DROP INDEX IF EXISTS public.idx_tokens_deleted_at_4 CASCADE;

-- Table: public.likes

DROP TABLE IF EXISTS public.likes CASCADE;

-- Index: idx_tokens_deleted_at_2

DROP INDEX IF EXISTS public.idx_tokens_deleted_at_2 CASCADE;

-- Table: public.forward_messages

DROP TABLE IF EXISTS public.forward_messages CASCADE;

-- Index: idx_message_files_deleted_at_1

DROP INDEX IF EXISTS public.idx_message_files_deleted_at_1 CASCADE;

-- Table: public.files

DROP TABLE IF EXISTS public.files CASCADE;

-- Index: idx_files_deleted_at

DROP INDEX IF EXISTS public.idx_files_deleted_at CASCADE;

-- Table: public.chats

DROP TABLE IF EXISTS public.chats CASCADE;

-- Index: idx_chats_deleted_at

DROP INDEX IF EXISTS public.idx_chats_deleted_at CASCADE;

-- Table: public.calls

DROP TABLE IF EXISTS public.calls CASCADE;

-- Index: idx_calls_deleted_at

DROP INDEX IF EXISTS public.idx_calls_deleted_at CASCADE;

-- Table: public.branches

DROP TABLE IF EXISTS public.branches CASCADE;

-- Index: idx_Branch_1

DROP INDEX IF EXISTS public."idx_Branch_1" CASCADE;

-- Table: public.activitys

DROP TABLE IF EXISTS public.activitys CASCADE;

-- Index: idx_tokens_deleted_at_1

DROP INDEX IF EXISTS public.idx_tokens_deleted_at_1 CASCADE;