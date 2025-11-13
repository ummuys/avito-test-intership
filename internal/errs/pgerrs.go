package errs

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrPGNotFound             = errors.New("not_found")
	ErrPGDuplicate            = errors.New("duplicate")
	ErrPGForeignKey           = errors.New("foreign_key_violation")
	ErrPGInvalidInput         = errors.New("invalid_input")
	ErrPGConstraint           = errors.New("constraint_violation")
	ErrPGNullViolation        = errors.New("not_null_violation")
	ErrPGCheckViolation       = errors.New("check_violation")
	ErrPGExclusionViolation   = errors.New("exclusion_violation")
	ErrPGSerializationFailure = errors.New("serialization_failure")
	ErrPGDeadlock             = errors.New("deadlock_detected")

	ErrPGNumericOutOfRange  = errors.New("numeric_value_out_of_range")
	ErrPGInvalidTextFormat  = errors.New("invalid_text_representation")
	ErrPGStringTooLong      = errors.New("string_data_right_truncation")
	ErrPGInvalidDatetime    = errors.New("invalid_datetime_format")
	ErrPGDivisionByZero     = errors.New("division_by_zero")
	ErrPGUntranslatableChar = errors.New("untranslatable_character")

	ErrPGInsufficientPrivilege = errors.New("insufficient_privilege")
	ErrPGInvalidPassword       = errors.New("invalid_password")
	ErrPGUnauthorized          = errors.New("unauthorized")
	ErrPGUndefinedTable        = errors.New("undefined_table")
	ErrPGUndefinedColumn       = errors.New("undefined_column")

	ErrPGConnection     = errors.New("connection_failure")
	ErrPGTxState        = errors.New("invalid_transaction_state")
	ErrPGInFailedTx     = errors.New("in_failed_sql_transaction")
	ErrPGIdleTxTimeout  = errors.New("idle_in_transaction_timeout")
	ErrPGConnectionLost = errors.New("connection_lost")

	ErrPGOutOfMemory           = errors.New("out_of_memory")
	ErrPGDiskFull              = errors.New("disk_full")
	ErrPGTooManyConnections    = errors.New("too_many_connections")
	ErrPGSystemIO              = errors.New("io_error")
	ErrPGSystemInternal        = errors.New("system_error")
	ErrPGDataCorrupted         = errors.New("data_corrupted")
	ErrPGIndexCorrupted        = errors.New("index_corrupted")
	ErrPGConfigurationExceeded = errors.New("configuration_limit_exceeded")

	ErrPGSyntaxError      = errors.New("syntax_error")
	ErrPGDatatypeMismatch = errors.New("datatype_mismatch")
	ErrPGDuplicateObject  = errors.New("duplicate_object")
	ErrPGInvalidSchema    = errors.New("invalid_schema_definition")
	ErrPGInvalidTableDef  = errors.New("invalid_table_definition")

	ErrPGQueryCanceled   = errors.New("query_canceled")
	ErrPGAdminShutdown   = errors.New("admin_shutdown")
	ErrPGCrashShutdown   = errors.New("crash_shutdown")
	ErrPGDatabaseDropped = errors.New("database_dropped")

	ErrPGInternal = errors.New("internal_error")
)

func ParsePgErr(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrPGNotFound
	}
	if errors.Is(err, pgx.ErrTxClosed) {
		return ErrPGTxState
	}

	var ErrPG *pgconn.PgError
	if errors.As(err, &ErrPG) {
		switch ErrPG.Code {

		case "23505":
			return ErrPGDuplicate
		case "23503":
			return ErrPGForeignKey
		case "23502":
			return ErrPGNullViolation
		case "23514":
			return ErrPGCheckViolation
		case "23P01":
			return ErrPGExclusionViolation

		case "22P02":
			return ErrPGInvalidInput
		case "22003":
			return ErrPGNumericOutOfRange
		case "22001":
			return ErrPGStringTooLong
		case "22007":
			return ErrPGInvalidDatetime
		case "22012":
			return ErrPGDivisionByZero
		case "22P05":
			return ErrPGUntranslatableChar

		case "08006", "08001", "08003":
			return ErrPGConnection
		case "57P03":
			return ErrPGConnectionLost

		case "25000":
			return ErrPGTxState
		case "25P02":
			return ErrPGInFailedTx
		case "25P03":
			return ErrPGIdleTxTimeout

		case "40001":
			return ErrPGSerializationFailure
		case "40P01":
			return ErrPGDeadlock

		case "42601":
			return ErrPGSyntaxError
		case "42804":
			return ErrPGDatatypeMismatch
		case "42701", "42702", "42703", "42704":
			return ErrPGDuplicateObject
		case "42P06":
			return ErrPGInvalidSchema
		case "42P17":
			return ErrPGInvalidTableDef
		case "42501":
			return ErrPGInsufficientPrivilege
		case "28P01":
			return ErrPGInvalidPassword

		case "53100":
			return ErrPGDiskFull
		case "53200":
			return ErrPGOutOfMemory
		case "53300":
			return ErrPGTooManyConnections
		case "53400":
			return ErrPGConfigurationExceeded

		case "55P03":
			return ErrPGSystemIO
		case "55000":
			return ErrPGSystemInternal

		case "57014":
			return ErrPGQueryCanceled
		case "57P01":
			return ErrPGAdminShutdown
		case "57P02":
			return ErrPGCrashShutdown
		case "57P04":
			return ErrPGDatabaseDropped

		case "58030":
			return ErrPGSystemIO
		case "XX000":
			return ErrPGInternal
		case "XX001":
			return ErrPGDataCorrupted
		case "XX002":
			return ErrPGIndexCorrupted

		default:
			return err
		}
	}

	return err
}
