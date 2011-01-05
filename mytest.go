package main

import ( "encoding/base32")
/*
BASE32("f") = "MY======"
BASE32("fo") = "MZXQ===="
BASE32("foo") = "MZXW6==="
BASE32("foob") = "MZXW6YQ="
BASE32("fooba") = "MZXW6YTB"
BASE32("foobar") = "MZXW6YTBOI======"
*/


func main() {
        encode := make([]byte, 100)
        decode := make([]byte, 100)

        in := "f"
        base32.StdEncoding.Encode(encode, []byte(in))
        println("ENC:", in, " " , string(encode), " MY======", base32.StdEncoding.EncodedLen(1))
        base32.StdEncoding.Decode(decode, []byte("MY======"))
        println("DEC:", in, " " , string(decode), base32.StdEncoding.DecodedLen(len("MY======")))

        in = "fo"
        base32.StdEncoding.Encode(encode, []byte(in))
        println("ENC:", in, " " , string(encode), " MZXQ====", base32.StdEncoding.EncodedLen(2))
        base32.StdEncoding.Decode(decode, []byte("MZXQ===="))
        println("DEC:", in, " " , string(decode), base32.StdEncoding.DecodedLen(len("MZXQ====")))

        in = "foo"
        base32.StdEncoding.Encode(encode, []byte(in))
        println("ENC:", in, " " , string(encode), " MZXW6===", base32.StdEncoding.EncodedLen(3))
        base32.StdEncoding.Decode(decode, []byte("MZXW6==="))
        println(in, " " , string(decode))
        println("DEC:", in, " " , string(decode), base32.StdEncoding.DecodedLen(len("MZXW6===")))

        in = "foob"
        base32.StdEncoding.Encode(encode, []byte(in))
        println("ENC:", in, " " , string(encode), " MZXW6YQ=", base32.StdEncoding.EncodedLen(4))
        base32.StdEncoding.Decode(decode, []byte("MZXW6YQ="))
        println(in, " " , string(decode))
        println("DEC:", in, " " , string(decode), base32.StdEncoding.DecodedLen(len("MZXW6YQ=")))

        in = "fooba"
        base32.StdEncoding.Encode(encode, []byte(in))
        println("ENC:", in, " " , string(encode), " MZXW6YTB", base32.StdEncoding.EncodedLen(5))
        base32.StdEncoding.Decode(decode, []byte("MZXW6YTB"))
        println(in, " " , string(decode))
        println("DEC:", in, " " , string(decode), base32.StdEncoding.DecodedLen(len("MZXW6YTB")))

        in = "foobar"
        base32.StdEncoding.Encode(encode, []byte(in))
        println("ENC:", in, " " , string(encode), " MZXW6YTBOI======", base32.StdEncoding.EncodedLen(6))
        base32.StdEncoding.Decode(decode, []byte("MZXW6YTBOI======"))
        println(in, " " , string(decode))
        println("DEC:", in, " " , string(decode), base32.StdEncoding.DecodedLen(len("MZXW6YTBOI======")))

        in = "Twas brillig, and the slithy toves"
        base32.StdEncoding.Encode(encode, []byte(in))
        println("ENC:", in, " " , string(encode), " KR3WC4ZAMJZGS3DMNFTSYIDBNZSCA5DIMUQHG3DJORUHSIDUN53GK4Y=", base32.StdEncoding.EncodedLen(35))
        base32.StdEncoding.Decode(decode, []byte("KR3WC4ZAMJZGS3DMNFTSYIDBNZSCA5DIMUQHG3DJORUHSIDUN53GK4Y="))
        println(in, " " , string(decode))
        println("DEC:", in, " " , string(decode), base32.StdEncoding.DecodedLen(len("KR3WC4ZAMJZGS3DMNFTSYIDBNZSCA5DIMUQHG3DJORUHSIDUN53GK4Y=")))

}
