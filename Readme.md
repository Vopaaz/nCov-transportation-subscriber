# nCov-transportation-subscriber

新型冠状病毒患者同行**订阅**查询工具。现有工具只能对每一班次进行手动查询，每天查询确认输入比较麻烦。
此工具可以订阅航班/列车班次等，每次进行查询无需重复输入。

WuHan coronavirus transportation subscriber. The current tools only supports manual querying, where inputting the date and flight number can be troublesome.
This tool enables you to subscribe a flight/train. After that you will not need to input that again.


## Installation

Clone this repo and run `go build`, or download from release (for Windows 10 x64).

## Usage

Suppose the built executable is named `ncov.exe`.

### Subscribe a Trip

```bash
./ncov.exe -d DATE -n NUMBER
```

`DATE` should be in `MM-dd`, for example `01-12` OR `1-2`. Note that the year is defaulted at 2020.
`NUMBER` is the flight/train number.

After subscription, a check will be automatically performed.

### Check the Subscribed Trips

Simply run

```bash
./ncov.exe
```

### Example

```bash
$ ./ncov -d 1-26 -n TR134
[WARNING] 2020-01-26 TR134 IS REPORTED!
Check successful.

$ ./ncov -d 1-25 -n 8L9564
[WARNING] 2020-01-26 TR134 IS REPORTED!
[WARNING] 2020-01-25 8L9564 IS REPORTED!
Check successful.

$ ./ncov
[WARNING] 2020-01-26 TR134 IS REPORTED!
[WARNING] 2020-01-25 8L9564 IS REPORTED!
Check successful.

$ ./ncov -d 1-24 -n not-dangerous
[WARNING] 2020-01-26 TR134 IS REPORTED!
[WARNING] 2020-01-25 8L9564 IS REPORTED!
Check successful.
```

## Data Source

[新型冠状病毒感染的肺炎确诊患者同行程查询工具](https://h5.peopleapp.com/txcx/index.html),
whose API locates at https://2019ncov.nosugartech.com/data.json

