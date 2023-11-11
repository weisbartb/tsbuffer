package tsbuffer_test

import (
	"github.com/stretchr/testify/require"
	"github.com/weisbartb/tsbuffer"
	"testing"
)

func TestNew(t *testing.T) {
	b := tsbuffer.New()
	require.NotEmpty(t, b)
	require.Empty(t, b.String())
	require.NoError(t, b.Close())
}

func TestSyncBuffer_Bytes(t *testing.T) {
	b := tsbuffer.New()
	n, err := b.Write([]byte("Test"))
	require.Equal(t, 4, n)
	require.NoError(t, err)
	require.NotEmpty(t, b.Bytes())
	require.Equal(t, []byte("Test"), b.Bytes())
}

func TestSyncBuffer_Close(t *testing.T) {
	b := tsbuffer.New()
	require.NoError(t, b.Close())
	n, err := b.WriteString("test")
	require.Error(t, err)
	require.Equal(t, 0, n)

}

func TestSyncBuffer_Len(t *testing.T) {
	b := tsbuffer.New()
	_, err := b.WriteString("test")
	require.NoError(t, err)
	require.Equal(t, 4, b.Len())
}

func TestSyncBuffer_Read(t *testing.T) {
	b := tsbuffer.New()
	_, err := b.WriteString("test1test2test3")
	require.NoError(t, err)
	var buf = make([]byte, 5)
	n, err := b.Read(buf)
	require.NoError(t, err)
	require.Equal(t, 5, n)
	require.Equal(t, "test1", string(buf))
	n, err = b.Read(buf)
	require.NoError(t, err)
	require.Equal(t, 5, n)
	require.Equal(t, "test2", string(buf))
	n, err = b.Read(buf)
	require.NoError(t, err)
	require.Equal(t, 5, n)
	require.Equal(t, "test3", string(buf))
	n, err = b.Read(buf)
	require.Error(t, err)
	require.Equal(t, 0, n)
}

func TestSyncBuffer_String(t *testing.T) {
	b := tsbuffer.New()
	n, err := b.Write([]byte("Test"))
	require.Equal(t, 4, n)
	require.NoError(t, err)
	require.NotEmpty(t, b.Bytes())
	require.Equal(t, "Test", b.String())
}

func TestSyncBuffer_Truncate(t *testing.T) {
	b := tsbuffer.New()
	_, err := b.WriteString("test1test2test3")
	require.NoError(t, err)
	var buf = make([]byte, 5)
	b.Truncate(4)
	n, err := b.Read(buf)
	require.NoError(t, err)
	require.Equal(t, 4, n)
	require.Equal(t, "test\x00", string(buf))
}
