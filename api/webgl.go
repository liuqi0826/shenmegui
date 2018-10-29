package api

import "syscall/js"

const (
	GL_DEPTH_BUFFER_BIT   = uint32(0x00000100)
	GL_STENCIL_BUFFER_BIT = uint32(0x00000400)
	GL_COLOR_BUFFER_BIT   = uint32(0x00004000)

	GL_POINTS         = uint32(0x0000)
	GL_LINES          = uint32(0x0001)
	GL_LINE_LOOP      = uint32(0x0002)
	GL_LINE_STRIP     = uint32(0x0003)
	GL_TRIANGLES      = uint32(0x0004)
	GL_TRIANGLE_STRIP = uint32(0x0005)
	GL_TRIANGLE_FAN   = uint32(0x0006)

	GL_ZERO                     = uint32(0)
	GL_ONE                      = uint32(1)
	GL_SRC_COLOR                = uint32(0x0300)
	GL_ONE_MINUS_SRC_COLOR      = uint32(0x0301)
	GL_SRC_ALPHA                = uint32(0x0302)
	GL_ONE_MINUS_SRC_ALPHA      = uint32(0x0303)
	GL_DST_ALPHA                = uint32(0x0304)
	GL_ONE_MINUS_DST_ALPHA      = uint32(0x0305)
	GL_DST_COLOR                = uint32(0x0306)
	GL_ONE_MINUS_DST_COLOR      = uint32(0x0307)
	GL_SRC_ALPHA_SATURATE       = uint32(0x0308)
	GL_CONSTANT_COLOR           = uint32(0x8001)
	GL_ONE_MINUS_CONSTANT_COLOR = uint32(0x8002)
	GL_CONSTANT_ALPHA           = uint32(0x8003)
	GL_ONE_MINUS_CONSTANT_ALPHA = uint32(0x8004)

	GL_FUNC_ADD              = uint32(0x8006)
	GL_FUNC_SUBSTRACT        = uint32(0x800A)
	GL_FUNC_REVERSE_SUBTRACT = uint32(0x800B)

	GL_BLEND_EQUATION                   = uint32(0x8009)
	GL_BLEND_EQUATION_RGB               = uint32(0x8009)
	GL_BLEND_EQUATION_ALPHA             = uint32(0x883D)
	GL_BLEND_DST_RGB                    = uint32(0x80C8)
	GL_BLEND_SRC_RGB                    = uint32(0x80C9)
	GL_BLEND_DST_ALPHA                  = uint32(0x80CA)
	GL_BLEND_SRC_ALPHA                  = uint32(0x80CB)
	GL_BLEND_COLOR                      = uint32(0x8005)
	GL_ARRAY_BUFFER_BINDING             = uint32(0x8894)
	GL_ELEMENT_ARRAY_BUFFER_BINDING     = uint32(0x8895)
	GL_LINE_WIDTH                       = uint32(0x0B21)
	GL_ALIASED_POINT_SIZE_RANGE         = uint32(0x846D)
	GL_ALIASED_LINE_WIDTH_RANGE         = uint32(0x846E)
	GL_CULL_FACE_MODE                   = uint32(0x0B45)
	GL_FRONT_FACE                       = uint32(0x0B46)
	GL_DEPTH_RANGE                      = uint32(0x0B70)
	GL_DEPTH_WRITEMASK                  = uint32(0x0B72)
	GL_DEPTH_CLEAR_VALUE                = uint32(0x0B73)
	GL_DEPTH_FUNC                       = uint32(0x0B74)
	GL_STENCIL_CLEAR_VALUE              = uint32(0x0B91)
	GL_STENCIL_FUNC                     = uint32(0x0B92)
	GL_STENCIL_FAIL                     = uint32(0x0B94)
	GL_STENCIL_PASS_DEPTH_FAIL          = uint32(0x0B95)
	GL_STENCIL_PASS_DEPTH_PASS          = uint32(0x0B96)
	GL_STENCIL_REF                      = uint32(0x0B97)
	GL_STENCIL_VALUE_MASK               = uint32(0x0B93)
	GL_STENCIL_WRITEMASK                = uint32(0x0B98)
	GL_STENCIL_BACK_FUNC                = uint32(0x8800)
	GL_STENCIL_BACK_FAIL                = uint32(0x8801)
	GL_STENCIL_BACK_PASS_DEPTH_FAIL     = uint32(0x8802)
	GL_STENCIL_BACK_PASS_DEPTH_PASS     = uint32(0x8803)
	GL_STENCIL_BACK_REF                 = uint32(0x8CA3)
	GL_STENCIL_BACK_VALUE_MASK          = uint32(0x8CA4)
	GL_STENCIL_BACK_WRITEMASK           = uint32(0x8CA5)
	GL_VIEWPORT                         = uint32(0x0BA2)
	GL_SCISSOR_BOX                      = uint32(0x0C10)
	GL_COLOR_CLEAR_VALUE                = uint32(0x0C22)
	GL_COLOR_WRITEMASK                  = uint32(0x0C23)
	GL_UNPACK_ALIGNMENT                 = uint32(0x0CF5)
	GL_PACK_ALIGNMENT                   = uint32(0x0D05)
	GL_MAX_TEXTURE_SIZE                 = uint32(0x0D33)
	GL_MAX_VIEWPORT_DIMS                = uint32(0x0D3A)
	GL_SUBPIXEL_BITS                    = uint32(0x0D50)
	GL_RED_BITS                         = uint32(0x0D52)
	GL_GREEN_BITS                       = uint32(0x0D53)
	GL_BLUE_BITS                        = uint32(0x0D54)
	GL_ALPHA_BITS                       = uint32(0x0D55)
	GL_DEPTH_BITS                       = uint32(0x0D56)
	GL_STENCIL_BITS                     = uint32(0x0D57)
	GL_POLYGON_OFFSET_UNITS             = uint32(0x2A00)
	GL_POLYGON_OFFSET_FACTOR            = uint32(0x8038)
	GL_TEXTURE_BINDING_2D               = uint32(0x8069)
	GL_SAMPLE_BUFFERS                   = uint32(0x80A8)
	GL_SAMPLES                          = uint32(0x80A9)
	GL_SAMPLE_COVERAGE_VALUE            = uint32(0x80AA)
	GL_SAMPLE_COVERAGE_INVERT           = uint32(0x80AB)
	GL_COMPRESSED_TEXTURE_FORMATS       = uint32(0x86A3)
	GL_VENDOR                           = uint32(0x1F00)
	GL_RENDERER                         = uint32(0x1F01)
	GL_VERSION                          = uint32(0x1F02)
	GL_IMPLEMENTATION_COLOR_READ_TYPE   = uint32(0x8B9A)
	GL_IMPLEMENTATION_COLOR_READ_FORMAT = uint32(0x8B9B)
	GL_BROWSER_DEFAULT_WEBGL            = uint32(0x9244)

	GL_STATIC_DRAW          = uint32(0x88E4)
	GL_STREAM_DRAW          = uint32(0x88E0)
	GL_DYNAMIC_DRAW         = uint32(0x88E8)
	GL_ARRAY_BUFFER         = uint32(0x8892)
	GL_ELEMENT_ARRAY_BUFFER = uint32(0x8893)
	GL_BUFFER_SIZE          = uint32(0x8764)
	GL_BUFFER_USAGE         = uint32(0x8765)

	GL_CURRENT_VERTEX_ATTRIB              = uint32(0x8626)
	GL_VERTEX_ATTRIB_ARRAY_ENABLED        = uint32(0x8622)
	GL_VERTEX_ATTRIB_ARRAY_SIZE           = uint32(0x8623)
	GL_VERTEX_ATTRIB_ARRAY_STRIDE         = uint32(0x8624)
	GL_VERTEX_ATTRIB_ARRAY_TYPE           = uint32(0x8625)
	GL_VERTEX_ATTRIB_ARRAY_NORMALIZED     = uint32(0x886A)
	GL_VERTEX_ATTRIB_ARRAY_POINTER        = uint32(0x8645)
	GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING = uint32(0x889F)

	GL_CULL_FACE      = uint32(0x0B44)
	GL_FRONT          = uint32(0x0404)
	GL_BACK           = uint32(0x0405)
	GL_FRONT_AND_BACK = uint32(0x0408)

	GL_BLEND                    = uint32(0x0BE2)
	GL_DEPTH_TEST               = uint32(0x0B71)
	GL_DITHER                   = uint32(0x0BD0)
	GL_POLYGON_OFFSET_FILL      = uint32(0x8037)
	GL_SAMPLE_ALPHA_TO_COVERAGE = uint32(0x809E)
	GL_SAMPLE_COVERAGE          = uint32(0x80A0)
	GL_SCISSOR_TEST             = uint32(0x0C11)
	GL_STENCIL_TEST             = uint32(0x0B90)

	GL_NO_ERROR           = uint32(0)
	GL_INVALID_ENUM       = uint32(0x0500)
	GL_INVALID_VALUE      = uint32(0x0501)
	GL_INVALID_OPERATION  = uint32(0x0502)
	GL_OUT_OF_MEMORY      = uint32(0x0505)
	GL_CONTEXT_LOST_WEBGL = uint32(0x9242)

	GL_CW  = uint32(0x0900)
	GL_CCW = uint32(0x0901)

	GL_DONT_CARE            = uint32(0x1100)
	GL_FASTEST              = uint32(0x1101)
	GL_NICEST               = uint32(0x1102)
	GL_GENERATE_MIPMAP_HINT = uint32(0x8192)

	GL_BYTE           = uint32(0x1400)
	GL_UNSIGNED_BYTE  = uint32(0x1401)
	GL_SHORT          = uint32(0x1402)
	GL_UNSIGNED_SHORT = uint32(0x1403)
	GL_INT            = uint32(0x1404)
	GL_UNSIGNED_INT   = uint32(0x1405)
	GL_FLOAT          = uint32(0x1406)

	GL_DEPTH_COMPONENT = uint32(0x1902)
	GL_ALPHA           = uint32(0x1906)
	GL_RGB             = uint32(0x1907)
	GL_RGBA            = uint32(0x1908)
	GL_LUMINANCE       = uint32(0x1909)
	GL_LUMINANCE_ALPHA = uint32(0x190A)

	GL_UNSIGNED_SHORT_4_4_4_4 = uint32(0x8033)
	GL_UNSIGNED_SHORT_5_5_5_1 = uint32(0x8034)
	GL_UNSIGNED_SHORT_5_6_5   = uint32(0x8363)

	GL_FRAGMENT_SHADER                  = uint32(0x8B30)
	GL_VERTEX_SHADER                    = uint32(0x8B31)
	GL_COMPILE_STATUS                   = uint32(0x8B81)
	GL_DELETE_STATUS                    = uint32(0x8B80)
	GL_LINK_STATUS                      = uint32(0x8B82)
	GL_VALIDATE_STATUS                  = uint32(0x8B83)
	GL_ATTACHED_SHADERS                 = uint32(0x8B85)
	GL_ACTIVE_ATTRIBUTES                = uint32(0x8B89)
	GL_ACTIVE_UNIFORMS                  = uint32(0x8B86)
	GL_MAX_VERTEX_ATTRIBS               = uint32(0x8869)
	GL_MAX_VERTEX_UNIFORM_VECTORS       = uint32(0x8DFB)
	GL_MAX_VARYING_VECTORS              = uint32(0x8DFC)
	GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS = uint32(0x8B4D)
	GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS   = uint32(0x8B4C)
	GL_MAX_TEXTURE_IMAGE_UNITS          = uint32(0x8872)
	GL_MAX_FRAGMENT_UNIFORM_VECTORS     = uint32(0x8DFD)
	GL_SHADER_TYPE                      = uint32(0x8B4F)
	GL_SHADING_LANGUAGE_VERSION         = uint32(0x8B8C)
	GL_CURRENT_PROGRAM                  = uint32(0x8B8D)

	GL_NEVER    = uint32(0x0200)
	GL_ALWAYS   = uint32(0x0207)
	GL_LESS     = uint32(0x0201)
	GL_EQUAL    = uint32(0x0202)
	GL_LEQUAL   = uint32(0x0203)
	GL_GREATER  = uint32(0x0204)
	GL_GEQUAL   = uint32(0x0206)
	GL_NOTEQUAL = uint32(0x0205)

	GL_KEEP      = uint32(0x1E00)
	GL_REPLACE   = uint32(0x1E01)
	GL_INCR      = uint32(0x1E02)
	GL_DECR      = uint32(0x1E03)
	GL_INVERT    = uint32(0x150A)
	GL_INCR_WRAP = uint32(0x8507)
	GL_DECR_WRAP = uint32(0x8508)

	GL_NEAREST                     = uint32(0x2600)
	GL_LINEAR                      = uint32(0x2601)
	GL_NEAREST_MIPMAP_NEAREST      = uint32(0x2700)
	GL_LINEAR_MIPMAP_NEAREST       = uint32(0x2701)
	GL_NEAREST_MIPMAP_LINEAR       = uint32(0x2702)
	GL_LINEAR_MIPMAP_LINEAR        = uint32(0x2703)
	GL_TEXTURE_MAG_FILTER          = uint32(0x2800)
	GL_TEXTURE_MIN_FILTER          = uint32(0x2801)
	GL_TEXTURE_WRAP_S              = uint32(0x2802)
	GL_TEXTURE_WRAP_T              = uint32(0x2803)
	GL_TEXTURE_2D                  = uint32(0x0DE1)
	GL_TEXTURE                     = uint32(0x1702)
	GL_TEXTURE_CUBE_MAP            = uint32(0x8513)
	GL_TEXTURE_BINDING_CUBE_MAP    = uint32(0x8514)
	GL_TEXTURE_CUBE_MAP_POSITIVE_X = uint32(0x8515)
	GL_TEXTURE_CUBE_MAP_NEGATIVE_X = uint32(0x8516)
	GL_TEXTURE_CUBE_MAP_POSITIVE_Y = uint32(0x8517)
	GL_TEXTURE_CUBE_MAP_NEGATIVE_Y = uint32(0x8518)
	GL_TEXTURE_CUBE_MAP_POSITIVE_Z = uint32(0x8519)
	GL_TEXTURE_CUBE_MAP_NEGATIVE_Z = uint32(0x851A)
	GL_MAX_CUBE_MAP_TEXTURE_SIZE   = uint32(0x851C)
	GL_TEXTURE0                    = uint32(0x84C0)
	GL_TEXTURE1                    = uint32(0x09C1)
	GL_TEXTURE2                    = uint32(0x09C2)
	GL_TEXTURE3                    = uint32(0x09C3)
	GL_TEXTURE4                    = uint32(0x09C4)
	GL_TEXTURE5                    = uint32(0x09C5)
	GL_TEXTURE6                    = uint32(0x09C6)
	GL_TEXTURE7                    = uint32(0x09C7)
	GL_TEXTURE8                    = uint32(0x09C8)
	GL_TEXTURE9                    = uint32(0x09C9)
	GL_TEXTURE10                   = uint32(0x09CA)
	GL_TEXTURE11                   = uint32(0x09CB)
	GL_TEXTURE12                   = uint32(0x09CC)
	GL_TEXTURE13                   = uint32(0x09CD)
	GL_TEXTURE14                   = uint32(0x09CE)
	GL_TEXTURE15                   = uint32(0x09CF)
	GL_TEXTURE16                   = uint32(0x09D0)
	GL_TEXTURE17                   = uint32(0x09D1)
	GL_TEXTURE18                   = uint32(0x09D2)
	GL_TEXTURE19                   = uint32(0x09D3)
	GL_TEXTURE20                   = uint32(0x09D4)
	GL_TEXTURE21                   = uint32(0x09D5)
	GL_TEXTURE22                   = uint32(0x09D6)
	GL_TEXTURE23                   = uint32(0x09D7)
	GL_TEXTURE24                   = uint32(0x09D8)
	GL_TEXTURE25                   = uint32(0x09D9)
	GL_TEXTURE26                   = uint32(0x09DA)
	GL_TEXTURE27                   = uint32(0x09DB)
	GL_TEXTURE28                   = uint32(0x09DC)
	GL_TEXTURE29                   = uint32(0x09DD)
	GL_TEXTURE30                   = uint32(0x09DE)
	GL_TEXTURE31                   = uint32(0x84DF)
	GL_ACTIVE_TEXTURE              = uint32(0x84E0)
	GL_REPEAT                      = uint32(0x2901)
	GL_CLAMP_TO_EDGE               = uint32(0x812F)
	GL_MIRRORED_REPEAT             = uint32(0x8370)

	GL_FLOAT_VEC2   = uint32(0x8B50)
	GL_FLOAT_VEC3   = uint32(0x8B51)
	GL_FLOAT_VEC4   = uint32(0x8B52)
	GL_INT_VEC2     = uint32(0x8B53)
	GL_INT_VEC3     = uint32(0x8B54)
	GL_INT_VEC4     = uint32(0x8B55)
	GL_BOOL         = uint32(0x8B56)
	GL_BOOL_VEC2    = uint32(0x8B57)
	GL_BOOL_VEC3    = uint32(0x8B58)
	GL_BOOL_VEC4    = uint32(0x8B59)
	GL_FLOAT_MAT2   = uint32(0x8B5A)
	GL_FLOAT_MAT3   = uint32(0x8B5B)
	GL_FLOAT_MAT4   = uint32(0x8B5C)
	GL_SAMPLER_2D   = uint32(0x8B5E)
	GL_SAMPLER_CUBE = uint32(0x8B60)

	GL_LOW_FLOAT    = uint32(0x8DF0)
	GL_MEDIUM_FLOAT = uint32(0x8DF1)
	GL_HIGH_FLOAT   = uint32(0x8DF2)
	GL_LOW_INT      = uint32(0x8DF3)
	GL_MEDIUM_INT   = uint32(0x8DF4)
	GL_HIGH_INT     = uint32(0x8DF5)

	GL_FRAMEBUFFER                                  = uint32(0x8D40)
	GL_RENDERBUFFER                                 = uint32(0x8D41)
	GL_RGBA4                                        = uint32(0x8056)
	GL_RGB5_A1                                      = uint32(0x8057)
	GL_RGB565                                       = uint32(0x8D62)
	GL_DEPTH_COMPONENT16                            = uint32(0x81A5)
	GL_STENCIL_INDEX                                = uint32(0x1901)
	GL_STENCIL_INDEX8                               = uint32(0x8D48)
	GL_RENDERBUFFER_WIDTH                           = uint32(0x8D42)
	GL_RENDERBUFFER_HEIGHT                          = uint32(0x8D43)
	GL_RENDERBUFFER_INTERNAL_FORMAT                 = uint32(0x8D44)
	GL_RENDERBUFFER_RED_SIZE                        = uint32(0x8D50)
	GL_RENDERBUFFER_GREEN_SIZE                      = uint32(0x8D51)
	GL_RENDERBUFFER_BLUE_SIZE                       = uint32(0x8D52)
	GL_RENDERBUFFER_ALPHA_SIZE                      = uint32(0x8D53)
	GL_RENDERBUFFER_DEPTH_SIZE                      = uint32(0x8D54)
	GL_RENDERBUFFER_STENCIL_SIZE                    = uint32(0x8D55)
	GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           = uint32(0x8CD0)
	GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           = uint32(0x8CD1)
	GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         = uint32(0x8CD2)
	GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = uint32(0x8CD3)
	GL_COLOR_ATTACHMENT0                            = uint32(0x8CE0)
	GL_DEPTH_ATTACHMENT                             = uint32(0x8D00)
	GL_STENCIL_ATTACHMENT                           = uint32(0x8D20)
	GL_NONE                                         = uint32(0)
	GL_FRAMEBUFFER_COMPLETE                         = uint32(0x8CD5)
	GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT            = uint32(0x8CD6)
	GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT    = uint32(0x8CD7)
	GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS            = uint32(0x8CD9)
	GL_FRAMEBUFFER_UNSUPPORTED                      = uint32(0x8CDD)
	GL_FRAMEBUFFER_BINDING                          = uint32(0x8CA6)
	GL_RENDERBUFFER_BINDING                         = uint32(0x8CA7)
	GL_MAX_RENDERBUFFER_SIZE                        = uint32(0x84E8)
	GL_INVALID_FRAMEBUFFER_OPERATION                = uint32(0x0506)

	GL_UNPACK_FLIP_Y_WEBGL                = uint32(0x9240)
	GL_UNPACK_PREMULTIPLY_ALPHA_WEBGL     = uint32(0x9241)
	GL_UNPACK_COLORSPACE_CONVERSION_WEBGL = uint32(0x9243)

	GL_READ_BUFFER                     = uint32(0x0C02)
	GL_UNPACK_ROW_LENGTH               = uint32(0x0CF2)
	GL_UNPACK_SKIP_ROWS                = uint32(0x0CF3)
	GL_UNPACK_SKIP_PIXELS              = uint32(0x0CF4)
	GL_PACK_ROW_LENGTH                 = uint32(0x0D02)
	GL_PACK_SKIP_ROWS                  = uint32(0x0D03)
	GL_PACK_SKIP_PIXELS                = uint32(0x0D04)
	GL_TEXTURE_BINDING_3D              = uint32(0x806A)
	GL_UNPACK_SKIP_IMAGES              = uint32(0x806D)
	GL_UNPACK_IMAGE_HEIGHT             = uint32(0x806E)
	GL_MAX_3D_TEXTURE_SIZE             = uint32(0x8073)
	GL_MAX_ELEMENTS_VERTICES           = uint32(0x80E8)
	GL_MAX_ELEMENTS_INDICES            = uint32(0x80E9)
	GL_MAX_TEXTURE_LOD_BIAS            = uint32(0x84FD)
	GL_MAX_FRAGMENT_UNIFORM_COMPONENTS = uint32(0x8B49)
	GL_MAX_VERTEX_UNIFORM_COMPONENTS   = uint32(0x8B4A)
	GL_MAX_ARRAY_TEXTURE_LAYERS        = uint32(0x88FF)
	GL_MIN_PROGRAM_TEXEL_OFFSET        = uint32(0x8904)
	GL_MAX_PROGRAM_TEXEL_OFFSET        = uint32(0x8905)
	GL_MAX_VARYING_COMPONENTS          = uint32(0x8B4B)
	GL_FRAGMENT_SHADER_DERIVATIVE_HINT = uint32(0x8B8B)
	GL_RASTERIZER_DISCARD              = uint32(0x8C89)
	GL_VERTEX_ARRAY_BINDING            = uint32(0x85B5)
	GL_MAX_VERTEX_OUTPUT_COMPONENTS    = uint32(0x9122)
	GL_MAX_FRAGMENT_INPUT_COMPONENTS   = uint32(0x9125)
	GL_MAX_SERVER_WAIT_TIMEOUT         = uint32(0x9111)
	GL_MAX_ELEMENT_INDEX               = uint32(0x8D6B)

	GL_RED                      = uint32(0x1903)
	GL_RGB8                     = uint32(0x8051)
	GL_RGBA8                    = uint32(0x8058)
	GL_RGB10_A2                 = uint32(0x8059)
	GL_TEXTURE_3D               = uint32(0x806F)
	GL_TEXTURE_WRAP_R           = uint32(0x8072)
	GL_TEXTURE_MIN_LOD          = uint32(0x813A)
	GL_TEXTURE_MAX_LOD          = uint32(0x813B)
	GL_TEXTURE_BASE_LEVEL       = uint32(0x813C)
	GL_TEXTURE_MAX_LEVEL        = uint32(0x813D)
	GL_TEXTURE_COMPARE_MODE     = uint32(0x884C)
	GL_TEXTURE_COMPARE_FUNC     = uint32(0x884D)
	GL_SRGB                     = uint32(0x8C40)
	GL_SRGB8                    = uint32(0x8C41)
	GL_SRGB8_ALPHA8             = uint32(0x8C43)
	GL_COMPARE_REF_TO_TEXTURE   = uint32(0x884E)
	GL_RGBA32F                  = uint32(0x8814)
	GL_RGB32F                   = uint32(0x8815)
	GL_RGBA16F                  = uint32(0x881A)
	GL_RGB16F                   = uint32(0x881B)
	GL_TEXTURE_2D_ARRAY         = uint32(0x8C1A)
	GL_TEXTURE_BINDING_2D_ARRAY = uint32(0x8C1D)
	GL_R11F_G11F_B10F           = uint32(0x8C3A)
	GL_RGB9_E5                  = uint32(0x8C3D)
	GL_RGBA32UI                 = uint32(0x8D70)
	GL_RGB32UI                  = uint32(0x8D71)
	GL_RGBA16UI                 = uint32(0x8D76)
	GL_RGB16UI                  = uint32(0x8D77)
	GL_RGBA8UI                  = uint32(0x8D7C)
	GL_RGB8UI                   = uint32(0x8D7D)
	GL_RGBA32I                  = uint32(0x8D82)
	GL_RGB32I                   = uint32(0x8D83)
	GL_RGBA16I                  = uint32(0x8D88)
	GL_RGB16I                   = uint32(0x8D89)
	GL_RGBA8I                   = uint32(0x8D8E)
	GL_RGB8I                    = uint32(0x8D8F)
	GL_RED_INTEGER              = uint32(0x8D94)
	GL_RGB_INTEGER              = uint32(0x8D98)
	GL_RGBA_INTEGER             = uint32(0x8D99)
	GL_R8                       = uint32(0x8229)
	GL_RG8                      = uint32(0x822B)
	GL_R16F                     = uint32(0x822D)
	GL_R32F                     = uint32(0x822E)
	GL_RG16F                    = uint32(0x822F)
	GL_RG32F                    = uint32(0x8230)
	GL_R8I                      = uint32(0x8231)
	GL_R8UI                     = uint32(0x8232)
	GL_R16I                     = uint32(0x8233)
	GL_R16UI                    = uint32(0x8234)
	GL_R32I                     = uint32(0x8235)
	GL_R32UI                    = uint32(0x8236)
	GL_RG8I                     = uint32(0x8237)
	GL_RG8UI                    = uint32(0x8238)
	GL_RG16I                    = uint32(0x8239)
	GL_RG16UI                   = uint32(0x823A)
	GL_RG32I                    = uint32(0x823B)
	GL_RG32UI                   = uint32(0x823C)
	GL_R8_SNORM                 = uint32(0x8F94)
	GL_RG8_SNORM                = uint32(0x8F95)
	GL_RGB8_SNORM               = uint32(0x8F96)
	GL_RGBA8_SNORM              = uint32(0x8F97)
	GL_RGB10_A2UI               = uint32(0x906F)
	GL_TEXTURE_IMMUTABLE_FORMAT = uint32(0x912F)
	GL_TEXTURE_IMMUTABLE_LEVELS = uint32(0x82DF)

	GL_UNSIGNED_INT_2_10_10_10_REV    = uint32(0x8368)
	GL_UNSIGNED_INT_10F_11F_11F_REV   = uint32(0x8C3B)
	GL_UNSIGNED_INT_5_9_9_9_REV       = uint32(0x8C3E)
	GL_FLOAT_32_UNSIGNED_INT_24_8_REV = uint32(0x8DAD)
	GL_UNSIGNED_INT_24_8              = uint32(0x84FA)
	GL_HALF_FLOAT                     = uint32(0x140B)
	GL_RG                             = uint32(0x8227)
	GL_RG_INTEGER                     = uint32(0x8228)
	GL_INT_2_10_10_10_REV             = uint32(0x8D9F)

	GL_CURRENT_QUERY                   = uint32(0x8865)
	GL_QUERY_RESULT                    = uint32(0x8866)
	GL_QUERY_RESULT_AVAILABLE          = uint32(0x8867)
	GL_ANY_SAMPLES_PASSED              = uint32(0x8C2F)
	GL_ANY_SAMPLES_PASSED_CONSERVATIVE = uint32(0x8D6A)

	GL_MAX_DRAW_BUFFERS      = uint32(0x8D9F)
	GL_DRAW_BUFFER0          = uint32(0x8D9F)
	GL_DRAW_BUFFER1          = uint32(0x8D9F)
	GL_DRAW_BUFFER2          = uint32(0x8D9F)
	GL_DRAW_BUFFER3          = uint32(0x8D9F)
	GL_DRAW_BUFFER4          = uint32(0x8D9F)
	GL_DRAW_BUFFER5          = uint32(0x8D9F)
	GL_DRAW_BUFFER6          = uint32(0x8D9F)
	GL_DRAW_BUFFER7          = uint32(0x8D9F)
	GL_DRAW_BUFFER8          = uint32(0x8D9F)
	GL_DRAW_BUFFER9          = uint32(0x8D9F)
	GL_DRAW_BUFFER10         = uint32(0x8D9F)
	GL_DRAW_BUFFER11         = uint32(0x8D9F)
	GL_DRAW_BUFFER12         = uint32(0x8D9F)
	GL_DRAW_BUFFER13         = uint32(0x8D9F)
	GL_DRAW_BUFFER14         = uint32(0x8D9F)
	GL_DRAW_BUFFER15         = uint32(0x8D9F)
	GL_MAX_COLOR_ATTACHMENTS = uint32(0x8D9F)
	GL_COLOR_ATTACHMENT1     = uint32(0x8CE1)
	GL_COLOR_ATTACHMENT2     = uint32(0x8CE2)
	GL_COLOR_ATTACHMENT3     = uint32(0x8CE3)
	GL_COLOR_ATTACHMENT4     = uint32(0x8CE4)
	GL_COLOR_ATTACHMENT5     = uint32(0x8CE5)
	GL_COLOR_ATTACHMENT6     = uint32(0x8CE6)
	GL_COLOR_ATTACHMENT7     = uint32(0x8CE7)
	GL_COLOR_ATTACHMENT8     = uint32(0x8CE8)
	GL_COLOR_ATTACHMENT9     = uint32(0x8CE9)
	GL_COLOR_ATTACHMENT10    = uint32(0x8CEA)
	GL_COLOR_ATTACHMENT11    = uint32(0x8CEB)
	GL_COLOR_ATTACHMENT12    = uint32(0x8CEC)
	GL_COLOR_ATTACHMENT13    = uint32(0x8CED)
	GL_COLOR_ATTACHMENT14    = uint32(0x8CEE)
	GL_COLOR_ATTACHMENT15    = uint32(0x8CEF)

	GL_SAMPLER_3D                    = uint32(0x8B5F)
	GL_SAMPLER_2D_SHADOW             = uint32(0x8B62)
	GL_SAMPLER_2D_ARRAY              = uint32(0x8DC1)
	GL_SAMPLER_2D_ARRAY_SHADOW       = uint32(0x8DC4)
	GL_SAMPLER_CUBE_SHADOW           = uint32(0x8DC5)
	GL_INT_SAMPLER_2D                = uint32(0x8DCA)
	GL_INT_SAMPLER_3D                = uint32(0x8DCB)
	GL_INT_SAMPLER_CUBE              = uint32(0x8DCC)
	GL_INT_SAMPLER_2D_ARRAY          = uint32(0x8DCF)
	GL_UNSIGNED_INT_SAMPLER_2D       = uint32(0x8DD2)
	GL_UNSIGNED_INT_SAMPLER_3D       = uint32(0x8DD3)
	GL_UNSIGNED_INT_SAMPLER_CUBE     = uint32(0x8DD4)
	GL_UNSIGNED_INT_SAMPLER_2D_ARRAY = uint32(0x8DD7)
	GL_MAX_SAMPLES                   = uint32(0x8D57)
	GL_SAMPLER_BINDING               = uint32(0x8919)

	GL_PIXEL_PACK_BUFFER           = uint32(0x88EB)
	GL_PIXEL_UNPACK_BUFFER         = uint32(0x88EC)
	GL_PIXEL_PACK_BUFFER_BINDING   = uint32(0x88ED)
	GL_PIXEL_UNPACK_BUFFER_BINDING = uint32(0x88EF)
	GL_COPY_READ_BUFFER            = uint32(0x8F36)
	GL_COPY_WRITE_BUFFER           = uint32(0x8F37)
	GL_COPY_READ_BUFFER_BINDING    = uint32(0x8F36)
	GL_COPY_WRITE_BUFFER_BINDING   = uint32(0x8F37)

	GL_FLOAT_MAT2x3        = uint32(0x8B65)
	GL_FLOAT_MAT2x4        = uint32(0x8B66)
	GL_FLOAT_MAT3x2        = uint32(0x8B67)
	GL_FLOAT_MAT3x4        = uint32(0x8B68)
	GL_FLOAT_MAT4x2        = uint32(0x8B69)
	GL_FLOAT_MAT4x3        = uint32(0x8B6A)
	GL_UNSIGNED_INT_VEC2   = uint32(0x8DC6)
	GL_UNSIGNED_INT_VEC3   = uint32(0x8DC7)
	GL_UNSIGNED_INT_VEC4   = uint32(0x8DC8)
	GL_UNSIGNED_NORMALIZED = uint32(0x8C17)
	GL_SIGNED_NORMALIZED   = uint32(0x8F9C)

	GL_VERTEX_ATTRIB_ARRAY_INTEGER = uint32(0x88FD)
	GL_VERTEX_ATTRIB_ARRAY_DIVISOR = uint32(0x88FE)

	GL_TRANSFORM_FEEDBACK_BUFFER_MODE                = uint32(0x8C7F)
	GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS    = uint32(0x8C80)
	GL_TRANSFORM_FEEDBACK_VARYINGS                   = uint32(0x8C83)
	GL_TRANSFORM_FEEDBACK_BUFFER_START               = uint32(0x8C84)
	GL_TRANSFORM_FEEDBACK_BUFFER_SIZE                = uint32(0x8C85)
	GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN         = uint32(0x8C88)
	GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS = uint32(0x8C8A)
	GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS       = uint32(0x8C8B)
	GL_INTERLEAVED_ATTRIBS                           = uint32(0x8C8C)
	GL_SEPARATE_ATTRIBS                              = uint32(0x8C8D)
	GL_TRANSFORM_FEEDBACK_BUFFER                     = uint32(0x8C8E)
	GL_TRANSFORM_FEEDBACK_BUFFER_BINDING             = uint32(0x8C8F)
	GL_TRANSFORM_FEEDBACK                            = uint32(0x8E22)
	GL_TRANSFORM_FEEDBACK_PAUSED                     = uint32(0x8E23)
	GL_TRANSFORM_FEEDBACK_ACTIVE                     = uint32(0x8E24)
	GL_TRANSFORM_FEEDBACK_BINDING                    = uint32(0x8E25)

	GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING = uint32(0x8210)
	GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE = uint32(0x8211)
	GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE       = uint32(0x8212)
	GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE     = uint32(0x8213)
	GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE      = uint32(0x8214)
	GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE     = uint32(0x8215)
	GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE     = uint32(0x8216)
	GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE   = uint32(0x8217)
	GL_FRAMEBUFFER_DEFAULT                   = uint32(0x8218)
	GL_DEPTH_STENCIL_ATTACHMENT              = uint32(0x821A)
	GL_DEPTH_STENCIL                         = uint32(0x84F9)
	GL_DEPTH24_STENCIL8                      = uint32(0x88F0)
	GL_DRAW_FRAMEBUFFER_BINDING              = uint32(0x8CA6)
	GL_READ_FRAMEBUFFER                      = uint32(0x8CA8)
	GL_DRAW_FRAMEBUFFER                      = uint32(0x8CA9)
	GL_READ_FRAMEBUFFER_BINDING              = uint32(0x8CAA)
	GL_RENDERBUFFER_SAMPLES                  = uint32(0x8CAB)
	GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER  = uint32(0x8CD4)
	GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE    = uint32(0x8D56)

	GL_UNIFORM_BUFFER                              = uint32(0x8A11)
	GL_UNIFORM_BUFFER_BINDING                      = uint32(0x8A28)
	GL_UNIFORM_BUFFER_START                        = uint32(0x8A29)
	GL_UNIFORM_BUFFER_SIZE                         = uint32(0x8A2A)
	GL_MAX_VERTEX_UNIFORM_BLOCKS                   = uint32(0x8A2B)
	GL_MAX_FRAGMENT_UNIFORM_BLOCKS                 = uint32(0x8A2D)
	GL_MAX_COMBINED_UNIFORM_BLOCKS                 = uint32(0x8A2E)
	GL_MAX_UNIFORM_BUFFER_BINDINGS                 = uint32(0x8A2F)
	GL_MAX_UNIFORM_BLOCK_SIZE                      = uint32(0x8A30)
	GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS      = uint32(0x8A31)
	GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS    = uint32(0x8A33)
	GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT             = uint32(0x8A34)
	GL_ACTIVE_UNIFORM_BLOCKS                       = uint32(0x8A36)
	GL_UNIFORM_TYPE                                = uint32(0x8A37)
	GL_UNIFORM_SIZE                                = uint32(0x8A38)
	GL_UNIFORM_BLOCK_INDEX                         = uint32(0x8A3A)
	GL_UNIFORM_OFFSET                              = uint32(0x8A3B)
	GL_UNIFORM_ARRAY_STRIDE                        = uint32(0x8A3C)
	GL_UNIFORM_MATRIX_STRIDE                       = uint32(0x8A3D)
	GL_UNIFORM_IS_ROW_MAJOR                        = uint32(0x8A3E)
	GL_UNIFORM_BLOCK_BINDING                       = uint32(0x8A3F)
	GL_UNIFORM_BLOCK_DATA_SIZE                     = uint32(0x8A40)
	GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS               = uint32(0x8A42)
	GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES        = uint32(0x8A43)
	GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER   = uint32(0x8A44)
	GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER = uint32(0x8A46)

	GL_OBJECT_TYPE                = uint32(0x9112)
	GL_SYNC_CONDITION             = uint32(0x9113)
	GL_SYNC_STATUS                = uint32(0x9114)
	GL_SYNC_FLAGS                 = uint32(0x9115)
	GL_SYNC_FENCE                 = uint32(0x9116)
	GL_SYNC_GPU_COMMANDS_COMPLETE = uint32(0x9117)
	GL_UNSIGNALED                 = uint32(0x9118)
	GL_SIGNALED                   = uint32(0x9119)
	GL_ALREADY_SIGNALED           = uint32(0x911A)
	GL_TIMEOUT_EXPIRED            = uint32(0x911B)
	GL_CONDITION_SATISFIED        = uint32(0x911C)
	GL_WAIT_FAILED                = uint32(0x911D)
	GL_SYNC_FLUSH_COMMANDS_BIT    = uint32(0x00000001)

	GL_COLOR                         = uint32(0x1800)
	GL_DEPTH                         = uint32(0x1801)
	GL_STENCIL                       = uint32(0x1802)
	GL_MIN                           = uint32(0x8007)
	GL_MAX                           = uint32(0x8008)
	GL_DEPTH_COMPONENT24             = uint32(0x81A6)
	GL_STREAM_READ                   = uint32(0x88E1)
	GL_STREAM_COPY                   = uint32(0x88E2)
	GL_STATIC_READ                   = uint32(0x88E5)
	GL_STATIC_COPY                   = uint32(0x88E6)
	GL_DYNAMIC_READ                  = uint32(0x88E9)
	GL_DYNAMIC_COPY                  = uint32(0x88EA)
	GL_DEPTH_COMPONENT32F            = uint32(0x8CAC)
	GL_DEPTH32F_STENCIL8             = uint32(0x8CAD)
	GL_INVALID_INDEX                 = uint32(0xFFFFFFFF)
	GL_TIMEOUT_IGNORED               = int32(-1)
	GL_MAX_CLIENT_WAIT_TIMEOUT_WEBGL = uint32(0x9247)

	GL_VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE = uint32(0x88FE)

	GL_UNMASKED_VENDOR_WEBGL   = uint32(0x9245)
	GL_UNMASKED_RENDERER_WEBGL = uint32(0x9246)

	GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT = uint32(0x84FF)
	GL_TEXTURE_MAX_ANISOTROPY_EXT     = uint32(0x84FE)

	GL_COMPRESSED_RGB_S3TC_DXT1_EXT  = uint32(0x83F0)
	GL_COMPRESSED_RGBA_S3TC_DXT1_EXT = uint32(0x83F1)
	GL_COMPRESSED_RGBA_S3TC_DXT3_EXT = uint32(0x83F2)
	GL_COMPRESSED_RGBA_S3TC_DXT5_EXT = uint32(0x83F3)

	GL_COMPRESSED_R11_EAC                        = uint32(0x9270)
	GL_COMPRESSED_SIGNED_R11_EAC                 = uint32(0x9271)
	GL_COMPRESSED_RG11_EAC                       = uint32(0x9272)
	GL_COMPRESSED_SIGNED_RG11_EAC                = uint32(0x9273)
	GL_COMPRESSED_RGB8_ETC2                      = uint32(0x9274)
	GL_COMPRESSED_RGBA8_ETC2_EAC                 = uint32(0x9275)
	GL_COMPRESSED_SRGB8_ETC2                     = uint32(0x9276)
	GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC          = uint32(0x9277)
	GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2  = uint32(0x9278)
	GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2 = uint32(0x9279)

	GL_COMPRESSED_RGB_PVRTC_4BPPV1_IMG  = uint32(0x8C00)
	GL_COMPRESSED_RGBA_PVRTC_4BPPV1_IMG = uint32(0x8C02)
	GL_COMPRESSED_RGB_PVRTC_2BPPV1_IMG  = uint32(0x8C01)
	GL_COMPRESSED_RGBA_PVRTC_2BPPV1_IMG = uint32(0x8C03)

	GL_COMPRESSED_RGB_ETC1_WEBGL = uint32(0x8D64)

	GL_COMPRESSED_RGB_ATC_WEBGL                     = uint32(0x8C92)
	GL_COMPRESSED_RGBA_ATC_EXPLICIT_ALPHA_WEBGL     = uint32(0x8C92)
	GL_COMPRESSED_RGBA_ATC_INTERPOLATED_ALPHA_WEBGL = uint32(0x87EE)

	GL_UNSIGNED_INT_24_8_WEBGL = uint32(0x84FA)

	GL_HALF_FLOAT_OES = uint32(0x8D61)

	GL_RGBA32F_EXT                               = uint32(0x8814)
	GL_RGB32F_EXT                                = uint32(0x8815)
	GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT = uint32(0x8211)
	GL_UNSIGNED_NORMALIZED_EXT                   = uint32(0x8C17)

	GL_MIN_EXT = uint32(0x8007)
	GL_MAX_EXT = uint32(0x8008)

	GL_SRGB_EXT                                  = uint32(0x8C40)
	GL_SRGB_ALPHA_EXT                            = uint32(0x8C42)
	GL_SRGB8_ALPHA8_EXT                          = uint32(0x8C43)
	GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING_EXT = uint32(0x8210)

	GL_FRAGMENT_SHADER_DERIVATIVE_HINT_OES = uint32(0x8B8B)

	GL_COLOR_ATTACHMENT0_WEBGL     = uint32(0x8CE0)
	GL_COLOR_ATTACHMENT1_WEBGL     = uint32(0x8CE1)
	GL_COLOR_ATTACHMENT2_WEBGL     = uint32(0x8CE2)
	GL_COLOR_ATTACHMENT3_WEBGL     = uint32(0x8CE3)
	GL_COLOR_ATTACHMENT4_WEBGL     = uint32(0x8CE4)
	GL_COLOR_ATTACHMENT5_WEBGL     = uint32(0x8CE5)
	GL_COLOR_ATTACHMENT6_WEBGL     = uint32(0x8CE6)
	GL_COLOR_ATTACHMENT7_WEBGL     = uint32(0x8CE7)
	GL_COLOR_ATTACHMENT8_WEBGL     = uint32(0x8CE8)
	GL_COLOR_ATTACHMENT9_WEBGL     = uint32(0x8CE9)
	GL_COLOR_ATTACHMENT10_WEBGL    = uint32(0x8CEA)
	GL_COLOR_ATTACHMENT11_WEBGL    = uint32(0x8CEB)
	GL_COLOR_ATTACHMENT12_WEBGL    = uint32(0x8CEC)
	GL_COLOR_ATTACHMENT13_WEBGL    = uint32(0x8CED)
	GL_COLOR_ATTACHMENT14_WEBGL    = uint32(0x8CEE)
	GL_COLOR_ATTACHMENT15_WEBGL    = uint32(0x8CEF)
	GL_DRAW_BUFFER0_WEBGL          = uint32(0x8825)
	GL_DRAW_BUFFER1_WEBGL          = uint32(0x8826)
	GL_DRAW_BUFFER2_WEBGL          = uint32(0x8827)
	GL_DRAW_BUFFER3_WEBGL          = uint32(0x8828)
	GL_DRAW_BUFFER4_WEBGL          = uint32(0x8829)
	GL_DRAW_BUFFER5_WEBGL          = uint32(0x882A)
	GL_DRAW_BUFFER6_WEBGL          = uint32(0x882B)
	GL_DRAW_BUFFER7_WEBGL          = uint32(0x882C)
	GL_DRAW_BUFFER8_WEBGL          = uint32(0x882D)
	GL_DRAW_BUFFER9_WEBGL          = uint32(0x882E)
	GL_DRAW_BUFFER10_WEBGL         = uint32(0x882F)
	GL_DRAW_BUFFER11_WEBGL         = uint32(0x8830)
	GL_DRAW_BUFFER12_WEBGL         = uint32(0x8831)
	GL_DRAW_BUFFER13_WEBGL         = uint32(0x8832)
	GL_DRAW_BUFFER14_WEBGL         = uint32(0x8833)
	GL_DRAW_BUFFER15_WEBGL         = uint32(0x8834)
	GL_MAX_COLOR_ATTACHMENTS_WEBGL = uint32(0x8CDF)
	GL_MAX_DRAW_BUFFERS_WEBGL      = uint32(0x8824)

	GL_VERTEX_ARRAY_BINDING_OES = uint32(0x85B5)

	GL_QUERY_COUNTER_BITS_EXT     = uint32(0x8864)
	GL_CURRENT_QUERY_EXT          = uint32(0x8865)
	GL_QUERY_RESULT_EXT           = uint32(0x8866)
	GL_QUERY_RESULT_AVAILABLE_EXT = uint32(0x8867)
	GL_TIME_ELAPSED_EXT           = uint32(0x88BF)
	GL_TIMESTAMP_EXT              = uint32(0x8E28)
	GL_GPU_DISJOINT_EXT           = uint32(0x8FBB)
)

type WebGL struct {
	context js.Value
}

func (this *WebGL) Setup(context js.Value) {
	this.context = context
}
func (this *WebGL) ActiveTexture(texture uint32) {
	this.context.Call("activeTexture", texture)
}
func (this *WebGL) AttachShader(program js.Value, shader js.Value) {
	this.context.Call("attachShader", program, shader)
}
func (this *WebGL) BindAttribLocation(program js.Value, index uint32, name string) {
	this.context.Call("bindAttribLocation", program, index, name)
}
func (this *WebGL) BindBuffer(target uint32, buffer js.Value) {
	this.context.Call("bindBuffer", target, buffer)
}
func (this *WebGL) BindFramebuffer(target uint32, framebuffer js.Value) {
	this.context.Call("bindFramebuffer", target, framebuffer)
}
func (this *WebGL) BindRenderbuffer(target uint32, renderbuffer js.Value) {
	this.context.Call("bindRenderbuffer", target, renderbuffer)
}
func (this *WebGL) BindTexture(target uint32, texture js.Value) {
	this.context.Call("bindTexture", target, texture)
}
func (this *WebGL) BlendColor(red float32, green float32, blue float32, alpha float32) {
	this.context.Call("blendColor", red, green, blue, alpha)
}
func (this *WebGL) BlendEquation(mode uint32) {
	this.context.Call("blendEquation", mode)
}
func (this *WebGL) BlendEquationSeparate(modeRGB uint32, modeAlpha uint32) {
	this.context.Call("blendEquationSeparate", modeRGB, modeAlpha)
}
func (this *WebGL) BlendFunc(sfactor uint32, dfactor uint32) {
	this.context.Call("blendFunc", sfactor, dfactor)
}
func (this *WebGL) BlendFuncSeparate(srcRGB uint32, dstRGB uint32, srcAlpha uint32, dstAlpha uint32) {
	this.context.Call("blendFuncSeparate", srcRGB, dstRGB, srcAlpha, dstAlpha)
}
func (this *WebGL) BufferData(target uint32, srcData interface{}, usage uint32) {
	valueTypeArray := js.TypedArrayOf(srcData)
	this.context.Call("bufferData", target, valueTypeArray, usage)
}
func (this *WebGL) BufferSubData(target uint32, dstByteOffset js.Value, srcData js.Value, srcOffset uint32, length uint32) {
	this.context.Call("bufferSubData", target, dstByteOffset, srcData, srcOffset, length)
}
func (this *WebGL) CheckFramebufferStatus(target uint32) {
	this.context.Call("checkFramebufferStatus", target)
}
func (this *WebGL) Clear(mask uint32) {
	this.context.Call("clear", mask)
}
func (this *WebGL) ClearColor(red float32, green float32, blue float32, alpha float32) {
	this.context.Call("clearColor", red, green, blue, alpha)
}
func (this *WebGL) ClearDepth(depth float32) {
	this.context.Call("clearDepth", depth)
}
func (this *WebGL) ClearStencil(s int32) {
	this.context.Call("clearStencil", s)
}
func (this *WebGL) ColorMask(red bool, green bool, blue bool, alpha bool) {
	this.context.Call("colorMask", red, green, blue, alpha)
}
func (this *WebGL) Commit() {
	this.context.Call("commit")
}
func (this *WebGL) CompileShader(shader js.Value) {
	this.context.Call("compileShader", shader)
}
func (this *WebGL) CompressedTexImage2D(target uint32, level int32, internalformat uint32, width int32, height int32, border int32, pixels interface{}) {
	valueTypeArray := js.TypedArrayOf(pixels)
	this.context.Call("compressedTexImage2D", target, level, internalformat, width, height, border, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) CompressedTexImage3D(target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, srcData interface{}, srcOffset int32, srcLengthOverride int32) {
	valueTypeArray := js.TypedArrayOf(srcData)
	this.context.Call("compressedTexImage3D", target, level, internalformat, width, height, depth, border, valueTypeArray, srcOffset, srcLengthOverride)
	valueTypeArray.Release()
}
func (this *WebGL) CompressedTexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, pixels interface{}) {
	valueTypeArray := js.TypedArrayOf(pixels)
	this.context.Call("compressedTexSubImage2D", target, level, xoffset, yoffset, width, height, format, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) CopyTexImage2D(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, height int32, border int32) {
	this.context.Call("copyTexImage2D", target, level, internalformat, x, y, width, height, border)
}
func (this *WebGL) CopyTexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	this.context.Call("copyTexSubImage2D", target, level, xoffset, yoffset, x, y, width, height)
}
func (this *WebGL) CreateBuffer() js.Value {
	WebGLBuffer := this.context.Call("createBuffer")
	return WebGLBuffer
}
func (this *WebGL) CreateFramebuffer() js.Value {
	WebGLFramebuffer := this.context.Call("createFramebuffer")
	return WebGLFramebuffer
}
func (this *WebGL) CreateProgram() js.Value {
	WebGLProgram := this.context.Call("createProgram")
	return WebGLProgram
}
func (this *WebGL) CreateRenderbuffer() js.Value {
	WebGLRenderbuffer := this.context.Call("createRenderbuffer")
	return WebGLRenderbuffer
}
func (this *WebGL) CreateShader(gltype uint32) js.Value {
	WebGLShader := this.context.Call("createShader", gltype)
	return WebGLShader
}
func (this *WebGL) CreateTexture() js.Value {
	WebGLTexture := this.context.Call("createTexture")
	return WebGLTexture
}
func (this *WebGL) CullFace(mode uint32) {
	this.context.Call("cullFace", mode)
}
func (this *WebGL) DeleteBuffer(buffer js.Value) {
	this.context.Call("deleteBuffer", buffer)
}
func (this *WebGL) DeleteFramebuffer(framebuffer js.Value) {
	this.context.Call("deleteFramebuffer", framebuffer)
}
func (this *WebGL) DeleteProgram(program js.Value) {
	this.context.Call("deleteProgram", program)
}
func (this *WebGL) DeleteRenderbuffer(renderbuffer js.Value) {
	this.context.Call("deleteRenderbuffer", renderbuffer)
}
func (this *WebGL) DeleteShader(shader js.Value) {
	this.context.Call("deleteShader", shader)
}
func (this *WebGL) DeleteTexture(texture js.Value) {
	this.context.Call("deleteTexture", texture)
}
func (this *WebGL) DepthFunc(fun uint32) {
	this.context.Call("depthFunc", fun)
}
func (this *WebGL) DepthMask(flag bool) {
	this.context.Call("depthMask", flag)
}
func (this *WebGL) DepthRange(zNear float32, zFar float32) {
	this.context.Call("depthRange", zNear, zFar)
}
func (this *WebGL) DetachShader(program js.Value, shader js.Value) {
	this.context.Call("detachShader", program, shader)
}
func (this *WebGL) Disable(cap uint32) {
	this.context.Call("disable", cap)
}
func (this *WebGL) DisableVertexAttribArray(index uint32) {
	this.context.Call("disableVertexAttribArray", index)
}
func (this *WebGL) DrawArrays(mode uint32, first int32, count int32) {
	this.context.Call("drawArrays", mode, first, count)
}
func (this *WebGL) DrawElements(mode uint32, count int32, gltype uint32, offset int32) {
	this.context.Call("drawElements", mode, count, gltype, offset)
}
func (this *WebGL) Enable(cap uint32) {
	this.context.Call("enable", cap)
}
func (this *WebGL) EnableVertexAttribArray(index uint32) {
	this.context.Call("enableVertexAttribArray", index)
}
func (this *WebGL) Finish() {
	this.context.Call("finish")
}
func (this *WebGL) Flush() {
	this.context.Call("flush")
}
func (this *WebGL) FramebufferRenderbuffer(target uint32, attachment uint32, renderbuffertarget uint32, renderbuffer js.Value) {
	this.context.Call("framebufferRenderbuffer", target, attachment, renderbuffertarget, renderbuffer)
}
func (this *WebGL) FramebufferTexture2D(target uint32, attachment uint32, textarget uint32, texture js.Value, level int32) {
	this.context.Call("framebufferTexture2D", target, attachment, textarget, texture, level)
}
func (this *WebGL) FrontFace(mode uint32) {
	this.context.Call("frontFace", mode)
}
func (this *WebGL) GenerateMipmap(target uint32) {
	this.context.Call("generateMipmap", target)
}
func (this *WebGL) GetActiveAttrib(program js.Value, index uint32) js.Value {
	WebGLActiveInfo := this.context.Call("getActiveAttrib", program, index)
	return WebGLActiveInfo
}
func (this *WebGL) GetActiveUniform(program js.Value, index uint32) js.Value {
	WebGLActiveInfo := this.context.Call("getActiveUniform", program, index)
	return WebGLActiveInfo
}
func (this *WebGL) GetAttachedShaders(program js.Value) interface{} {
	sequence := this.context.Call("getAttachedShaders", program)
	return sequence
}
func (this *WebGL) GetAttribLocation(program js.Value, name string) uint32 {
	return uint32(this.context.Call("getAttribLocation", program, name).Int())
}
func (this *WebGL) GetBufferParameter(target uint32, pname uint32) interface{} {
	return this.context.Call("getBufferParameter", target, pname)
}
func (this *WebGL) GetContextAttributes() {
	this.context.Call("getContextAttributes")
}
func (this *WebGL) GetError() uint32 {
	code := uint32(this.context.Call("getError").Int())
	return code
}
func (this *WebGL) GetExtension(name string) {
	this.context.Call("getExtension", name)
}
func (this *WebGL) GetFramebufferAttachmentParameter(target uint32, attachment uint32, pname uint32) interface{} {
	return this.context.Call("getFramebufferAttachmentParameter", target, attachment, pname)
}
func (this *WebGL) GetParameter(pname uint32) interface{} {
	return this.context.Call("getParameter", pname)
}
func (this *WebGL) GetProgramInfoLog(program js.Value) js.Value {
	return this.context.Call("getProgramInfoLog", program)
}
func (this *WebGL) GetProgramParameter(program js.Value, pname uint32) interface{} {
	return this.context.Call("getProgramParameter", program, pname)
}
func (this *WebGL) GetRenderbufferParameter(target uint32, pname uint32) interface{} {
	return this.context.Call("getRenderbufferParameter", target, pname)
}
func (this *WebGL) GetShaderInfoLog(shader js.Value) js.Value {
	return this.context.Call("getShaderInfoLog", shader)
}
func (this *WebGL) GetShaderParameter(shader js.Value, pname uint32) interface{} {
	return this.context.Call("getShaderParameter", shader, pname)
}
func (this *WebGL) GetShaderPrecisionFormat(shaderType uint32, precisionType uint32) js.Value {
	WebGLShaderPrecisionFormat := this.context.Call("getShaderPrecisionFormat", shaderType, precisionType)
	return WebGLShaderPrecisionFormat
}
func (this *WebGL) GetShaderSource(shader js.Value) string {
	return this.context.Call("getShaderSource", shader).String()
}
func (this *WebGL) GetSupportedExtensions() {
	this.context.Call("getSupportedExtensions")
}
func (this *WebGL) GetTexParameter(target uint32, pname uint32) interface{} {
	return this.context.Call("getTexParameter", target, pname)
}
func (this *WebGL) GetUniform(program js.Value, location js.Value) interface{} {
	return this.context.Call("getUniform", program, location)
}
func (this *WebGL) GetUniformLocation(program js.Value, name string) js.Value {
	WebGLUniformLocation := this.context.Call("getUniformLocation", program, name)
	return WebGLUniformLocation
}
func (this *WebGL) GetVertexAttrib(index uint32, pname uint32) interface{} {
	return this.context.Call("getVertexAttrib", index, pname)
}
func (this *WebGL) GetVertexAttribOffset(index uint32, pname uint32) int32 {
	return int32(this.context.Call("getVertexAttribOffset", index, pname).Int())
}
func (this *WebGL) Hint(target uint32, mode uint32) {
	this.context.Call("hint", target, mode)
}
func (this *WebGL) IsBuffer(buffer js.Value) bool {
	return this.context.Call("isBuffer", buffer).Bool()
}
func (this *WebGL) IsContextLost() bool {
	return this.context.Call("isContextLost").Bool()
}
func (this *WebGL) IsEnabled(cap uint32) bool {
	return this.context.Call("isEnabled", cap).Bool()
}
func (this *WebGL) IsFramebuffer(framebuffer js.Value) bool {
	return this.context.Call("isFramebuffer", framebuffer).Bool()
}
func (this *WebGL) IsProgram(program js.Value) bool {
	return this.context.Call("isProgram", program).Bool()
}
func (this *WebGL) IsRenderbuffer(renderbuffer js.Value) bool {
	return this.context.Call("isRenderbuffer", renderbuffer).Bool()
}
func (this *WebGL) IsShader(shader js.Value) bool {
	return this.context.Call("isShader", shader).Bool()
}
func (this *WebGL) IsTexture(texture js.Value) bool {
	return this.context.Call("isTexture", texture).Bool()
}
func (this *WebGL) LineWidth(width float32) {
	this.context.Call("lineWidth", width)
}
func (this *WebGL) LinkProgram(program js.Value) {
	this.context.Call("linkProgram", program)
}
func (this *WebGL) PixelStorei(pname uint32, param int32) {
	this.context.Call("pixelStorei", pname, param)
}
func (this *WebGL) PolygonOffset(factor float32, units float32) {
	this.context.Call("polygonOffset", factor, units)
}
func (this *WebGL) ReadPixels(x int32, y int32, width int32, height int32, format uint32, gltype uint32, pixels js.Value) {
	this.context.Call("readPixels", x, y, width, height, format, gltype, pixels)
}
func (this *WebGL) RenderbufferStorage(target uint32, internalFormat uint32, width int32, height int32) {
	this.context.Call("renderbufferStorage", target, internalFormat, width, height)
}
func (this *WebGL) SampleCoverage(value float32, invert bool) {
	this.context.Call("sampleCoverage", value, invert)
}
func (this *WebGL) Scissor(x int32, y int32, width int32, height int32) {
	this.context.Call("scissor", x, y, width, height)
}
func (this *WebGL) ShaderSource(shader js.Value, source string) {
	this.context.Call("shaderSource", shader, source)
}
func (this *WebGL) StencilFunc(fun uint32, ref int32, mask uint32) {
	this.context.Call("stencilFunc", fun, ref, mask)
}
func (this *WebGL) StencilFuncSeparate(face uint32, fun uint32, ref int32, mask uint32) {
	this.context.Call("stencilFuncSeparate", face, fun, ref, mask)
}
func (this *WebGL) StencilMask(mask uint32) {
	this.context.Call("stencilMask", mask)
}
func (this *WebGL) StencilMaskSeparate(face uint32, mask uint32) {
	this.context.Call("stencilMaskSeparate", face, mask)
}
func (this *WebGL) StencilOp(fail uint32, zfail uint32, zpass uint32) {
	this.context.Call("stencilOp", fail, zfail, zpass)
}
func (this *WebGL) StencilOpSeparate(face uint32, fail uint32, zfail uint32, zpass uint32) {
	this.context.Call("stencilOpSeparate", face, fail, zfail, zpass)
}
func (this *WebGL) TexImage2D(target uint32, level int32, internalformat uint32, width int32, height int32, border int32, format uint32, gltype uint32, pixels interface{}) {
	valueTypeArray := js.TypedArrayOf(pixels)
	this.context.Call("texImage2D", target, level, internalformat, width, height, border, format, gltype, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) TexParameterf(target uint32, pname uint32, param float32) {
	this.context.Call("texParameterf", target, pname, param)
}
func (this *WebGL) TexParameteri(target uint32, pname uint32, param uint32) {
	this.context.Call("texParameterf", target, pname, param)
}
func (this *WebGL) TexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, gltype uint32, pixels interface{}) {
	valueTypeArray := js.TypedArrayOf(pixels)
	this.context.Call("texSubImage2D", target, level, xoffset, yoffset, width, height, format, gltype, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) Uniform1f(location js.Value, v0 float32) {
	this.context.Call("uniform1f", location, v0)
}
func (this *WebGL) Uniform1fv(location js.Value, value []float32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("uniform1fv", location, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) Uniform1i(location js.Value, v0 int32) {
	this.context.Call("uniform1i", location, v0)
}
func (this *WebGL) Uniform1iv(location js.Value, value []int32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("uniform1iv", location, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) Uniform2f(location js.Value, v0 float32, v1 float32) {
	this.context.Call("uniform2f", location, v0, v1)
}
func (this *WebGL) Uniform2fv(location js.Value, value []float32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("uniform2fv", location, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) Uniform2i(location js.Value, v0 int32, v1 int32) {
	this.context.Call("uniform2i", location, v0, v1)
}
func (this *WebGL) Uniform2iv(location js.Value, value []int32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("uniform2iv", location, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) Uniform3f(location js.Value, v0 float32, v1 float32, v2 float32) {
	this.context.Call("uniform3f", location, v0, v1, v2)
}
func (this *WebGL) Uniform3fv(location js.Value, value []float32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("uniform3fv", location, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) Uniform3i(location js.Value, v0 int32, v1 int32, v2 int32) {
	this.context.Call("uniform3i", location, v0, v1, v2)
}
func (this *WebGL) Uniform3iv(location js.Value, value []int32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("uniform3iv", location, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) Uniform4f(location js.Value, v0 float32, v1 float32, v2 float32, v3 float32) {
	this.context.Call("uniform4f", location, v0, v1, v2, v3)
}
func (this *WebGL) Uniform4fv(location js.Value, value []float32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("uniform4fv", location, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) Uniform4i(location js.Value, v0 int32, v1 int32, v2 int32, v3 int32) {
	this.context.Call("uniform4i", location, v0, v1, v2, v3)
}
func (this *WebGL) Uniform4iv(location js.Value, value []int32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("uniform4iv", location, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) UniformMatrix2fv(location js.Value, transpose bool, value []float32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("uniformMatrix2fv", location, transpose, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) UniformMatrix3fv(location js.Value, transpose bool, value []float32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("uniformMatrix3fv", location, transpose, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) UniformMatrix4fv(location js.Value, transpose bool, value []float32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("uniformMatrix4fv", location, transpose, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) UseProgram(program js.Value) {
	this.context.Call("useProgram", program)
}
func (this *WebGL) ValidateProgram(program js.Value) {
	this.context.Call("validateProgram", program)
}
func (this *WebGL) VertexAttrib1f(index uint32, v0 float32) {
	this.context.Call("vertexAttrib1f", index, v0)
}
func (this *WebGL) VertexAttrib1fv(index uint32, value []float32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("vertexAttrib1fv", index, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) VertexAttrib2f(index uint32, v0 float32, v1 float32) {
	this.context.Call("vertexAttrib2f", index, v0, v1)
}
func (this *WebGL) VertexAttrib2fv(index uint32, value []float32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("vertexAttrib2fv", index, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) VertexAttrib3f(index uint32, v0 float32, v1 float32, v2 float32) {
	this.context.Call("vertexAttrib3f", index, v0, v1, v2)
}
func (this *WebGL) VertexAttrib3fv(index uint32, value []float32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("vertexAttrib3fv", index, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) VertexAttrib4f(index uint32, v0 float32, v1 float32, v2 float32, v3 float32) {
	this.context.Call("vertexAttrib4f", index, v0, v1, v2, v3)
}
func (this *WebGL) VertexAttrib4fv(index uint32, value []float32) {
	valueTypeArray := js.TypedArrayOf(value)
	this.context.Call("vertexAttrib4fv", index, valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebGL) VertexAttribPointer(index uint32, size int32, gltype uint32, normalized bool, stride int32, offset int32) {
	this.context.Call("vertexAttribPointer", index, size, gltype, normalized, stride, offset)
}
func (this *WebGL) Viewport(x int32, y int32, width int32, height int32) {
	this.context.Call("viewport", x, y, width, height)
}
