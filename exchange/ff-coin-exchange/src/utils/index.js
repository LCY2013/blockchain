function swapMerge(direction, entrustType) {
    return direction == "BUY"
                ? (entrustType == "OPEN"
                  ? "openLong"
                  : "closeShort")
                : (entrustType == "OPEN"
                ? "openShort"
                : "closeLong")
}

function swapSplit(way) {
    let re = {
        direction: "",
        entrustType: ""
    }
    switch(way) {
        case "openLong":
            re.direction = "BUY"
            re.entrustType = "OPEN"
            break;
        case "closeShort":
            re.direction = "BUY"
            re.entrustType = "CLOSE"
            break;
        case "openShort":
            re.direction = "SELL"
            re.entrustType = "OPEN"
            break;
        case "closeLong":
            re.direction = "SELL"
            re.entrustType = "CLOSE"
            break;
        default:
            re.direction = ""
            re.entrustType = ""
    }
    return re;
}

export {
    swapMerge,
    swapSplit
}