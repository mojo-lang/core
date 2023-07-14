package core

const MediaTypeTypeName = "MediaType"
const MediaTypeTypeFullName = "mojo.core.MediaType"

const (
	ApplicationJson              = "application/json"
	ApplicationPdf               = "application/pdf"
	ApplicationOctetStream       = "application/octet-stream"
	ApplicationWwwFormUrlencoded = "application/x-www-form-urlencoded"
	ApplicationXml               = "application/xml"
	ApplicationGeoJson           = "application/geo+json"
	ApplicationZip               = "application/zip"

	AudioAAC    = "audio/aac"
	AudioMP4    = "audio/mp4"
	AudioMPEG   = "audio/mpeg"
	AudioOgg    = "audio/ogg"
	AudioVorbis = "audio/vorbis"
	AudioWave   = "audio/wave"
	AudioWebm   = "audio/webm"

	FontTTF  = "font/ttf"
	FontOTF  = "font/otf"
	FontWOFF = "font/woff"

	ImageGif  = "image/gif"
	ImageJpeg = "image/jpeg"
	ImagePng  = "image/png"
	ImageSvg  = "image/svg+xml"
	ImageWebp = "image/webp"

	ModelGltfBinary = "model/gltf-binary"
	ModelGltfJson   = "model/gltf+json"
	ModelOBJ        = "model/obj"
	ModelU3D        = "model/u3d"

	MultipartFormData   = "multipart/form-data"
	MultipartByteRanges = "multipart/byteranges"

	TextPlain      = "text/plain"
	TextCsv        = "text/csv"
	TextCss        = "text/css"
	TextHtml       = "text/html"
	TextJavascript = "text/javascript"

	VideoMp4  = "video/mp4"
	VideoOgg  = "video/ogg"
	VideoWebm = "video/webm"
)

func (x *MediaType) IsSame(mediaType string) bool {
	if x != nil {
		return x.Type+"/"+x.Subtype == mediaType
	}
	return false
}
