const { csvToJson } = require("@nancyel/utils");
const fs = require("fs");
const path = require("path");
const { exit } = require("process");

const filePath = process.argv[2];
if (!filePath) {
  console.info("Please specify the path to a csv file");
  exit(1);
}

const data = csvToJson(filePath, {
  delimiter: "\n",
  escapeCommas: true,
});

const directoryName = "output";
const outputDirPath = path.join(process.cwd(), "../", directoryName);

// Check if the directory already exists
if (!fs.existsSync(outputDirPath)) {
  fs.mkdirSync(outputDirPath);
  console.log(`Directory '${outputDirPath}' created.`);
}

// Write output file to ../output directory
fs.writeFileSync(
  `${outputDirPath}/lines_${new Date().getTime()}.json`,
  JSON.stringify(data)
);
