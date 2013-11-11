uuid.go
=======

A simple uuid wrapper for golang on linux. 

Example:
//Create
id := uuid.New()
uuidString := id.String()

//Clone an uudi
cdata := id.Bytes()
nid := uuid.FromBytes(cdata)

