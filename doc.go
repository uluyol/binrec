/*
Package binrec provides utility functions for reading and writing binary data from files.
Data is prefixed with a variable-length encoded int (compatible with encoding/binary.PutUvarint
and Java protobuf's writeDelimitedTo method.
*/
package binrec
