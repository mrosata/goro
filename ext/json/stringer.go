// Code generated by "stringer -type=JsonError,JsonEncOpt -output stringer.go"; DO NOT EDIT.

package json

import "strconv"

const _JsonError_name = "ErrNoneErrDepthErrStateMismatchErrCtrlCharErrSyntaxErrUtf8ErrRecursionErrInfOrNanErrUnsupportedTypeErrInvalidPropNameErrUtf16"

var _JsonError_index = [...]uint8{0, 7, 15, 31, 42, 51, 58, 70, 81, 99, 117, 125}

func (i JsonError) String() string {
	if i < 0 || i >= JsonError(len(_JsonError_index)-1) {
		return "JsonError(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _JsonError_name[_JsonError_index[i]:_JsonError_index[i+1]]
}

const _JsonEncOpt_name = "HexTagHexAmpHexAposHexQuotForceObjectNumericCheckUnescapedSlashesPrettyPrintUnescapedUnicodePartialOutputOnErrorPreserveZeroFractionUnescapedEOL"

var _JsonEncOpt_map = map[JsonEncOpt]string{
	1:    _JsonEncOpt_name[0:6],
	2:    _JsonEncOpt_name[6:12],
	4:    _JsonEncOpt_name[12:19],
	8:    _JsonEncOpt_name[19:26],
	16:   _JsonEncOpt_name[26:37],
	32:   _JsonEncOpt_name[37:49],
	64:   _JsonEncOpt_name[49:65],
	128:  _JsonEncOpt_name[65:76],
	256:  _JsonEncOpt_name[76:92],
	512:  _JsonEncOpt_name[92:112],
	1024: _JsonEncOpt_name[112:132],
	2048: _JsonEncOpt_name[132:144],
}

func (i JsonEncOpt) String() string {
	if str, ok := _JsonEncOpt_map[i]; ok {
		return str
	}
	return "JsonEncOpt(" + strconv.FormatInt(int64(i), 10) + ")"
}
