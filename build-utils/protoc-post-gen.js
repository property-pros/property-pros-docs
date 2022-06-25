const fs = require("fs");
const path = require("path");

const protoDir = path.join(__dirname, "../proto");
const files = fs.readdirSync(
    protoDir,
    { withFileTypes: true }
);

files.forEach((file) => {
    if (
        path.extname(file.name) === ".js" &&
        file.name.indexOf("config_pb") !== -1
    ) {
        const filename = path.join(
            protoDir,
            file.name
        );

        const fileData = fs.readFileSync(filename);

        const newContent = Buffer.from("const proto = {\r\nconfig: {\r\nv1: {}\r\n}\r\n};");

        fs.writeFileSync(filename, Buffer.concat([newContent, fileData]), {
            flag: "w",
        });
    }
});