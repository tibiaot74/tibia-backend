def sexToBool(sex):
    return sex.upper() == "MALE"


def outfitToInt(outfit, sex):
    if sex:
        mapper = {
            "warrior": 134,
            "summoner": 133,
            "noble": 132,
            "knight": 131,
            "mage": 130,
            "hunter": 129,
            "citzen": 128,
        }
    else:
        mapper = {
            "warrior": 142,
            "summoner": 141,
            "noble": 140,
            "knight": 139,
            "mage": 138,
            "hunter": 137,
            "citzen": 136,
        }
    return mapper[outfit]


def intToOutfit(outfitCode, sex):
    if sex:
        mapper = {
            134: "warrior",
            133: "summoner",
            132: "noble",
            131: "knight",
            130: "mage",
            129: "hunter",
            128: "citzen",
        }
    else:
        mapper = {
            142: "warrior",
            141: "summoner",
            140: "noble",
            139: "knight",
            138: "mage",
            137: "hunter",
            136: "citzen",
        }
    return mapper[outfitCode]