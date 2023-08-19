module('MahjongMgr', package.seeall)

--[[
    用于检测通用型牌型(3n+2) 升级替换递归检测 可极大缩小检测时间
    外部调用
    MahjongMgr.Init:        初始化数据
    MahjongMgr.CheckWin     检测是否成胡
    @Time     : 2020/6/25 15:04
    @Author   : 王一冰
    @Software : Copyright (c) 2020 中青宝互动网络股份有限公司
]]

MahjongArray = MahjongArray or {
    -- 序数牌组合
    ordinalArray = {},
    ordinalEyeArray= {},

    -- 风牌组合
    windArray = {},
    windEyeArray = {},
}

-- WARNNING:最大金牌数 可根据需求调整
MaxGold = 8

-- WARNNING:13张手牌的为4 16张手牌的为5
LEVEL = 4

CheckFengOne = CheckFengOne or {}
CheckFengTwo = CheckFengTwo or {}
CheckOrdinalOne = CheckOrdinalOne or {}
CheckOrdinalTwo = CheckOrdinalTwo or {}

function Init()
    local sTime = os.msectime()

    -- 每次调用都清掉
    MahjongArray = {
        -- 序数牌组合
        ordinalArray = {},
        ordinalEyeArray= {},

        -- 风牌组合
        windArray = {},
        windEyeArray = {},
    }
    CheckFengOne = {}
    CheckFengTwo = {}
    CheckOrdinalOne = {}
    CheckOrdinalTwo = {}

    for i = 0, MaxGold do
        MahjongArray.ordinalArray[i] = {}
        MahjongArray.ordinalEyeArray[i] = {}

        MahjongArray.windArray[i] = {}
        MahjongArray.windEyeArray[i] = {}

        CheckOrdinalOne[i] = {}
        CheckOrdinalTwo[i] = {}

        CheckFengOne[i] = {}
        CheckFengTwo[i] = {}
    end

    -- 生成所有序数牌组合
    GenerateOrdinalArray()

    -- 生成所有带将序数牌组合
    GenerateOrdinalArrayWithEye()

    -- 生成所有风牌组合
    GenerateFengArray()

    -- 生成所有带将风牌组合
    GenerateFengArrayWithEye()

    -- 释放
    CheckFengOne = nil
    CheckFengTwo = nil
    CheckOrdinalOne = nil
    CheckOrdinalTwo = nil

    local eTime = os.msectime()
    unilight.debug("MahjongMgr.Init time:"..(eTime-sTime))
end

-------------------------------------------------------------------------------

function GenerateOrdinalArray()
    local cards = {0, 0, 0, 0, 0, 0, 0, 0, 0}
    GenerateOrdinalArraySub(cards, 1, false)
end

function GenerateOrdinalArrayWithEye()
    local cards = {0, 0, 0, 0, 0, 0, 0, 0, 0}
    for i = 1, 9 do
        cards[i] = 2
        ParseOrdinalArray(cards, true)
        GenerateOrdinalArraySub(cards, 1, true)
        cards[i] = 0
    end
end

function GenerateOrdinalArraySub(cards, level, eye)
    for i = 1, 16 do
        repeat
            if i <= 9 then
                if cards[i] > 3 then
                    break
                end
                cards[i] = cards[i] + 3
            elseif i <= 16 then
                local index = i - 9
                if cards[index] >= 4 or cards[index + 1] >= 4 or cards[index + 2] >= 4 then
                    break
                end
                cards[index] = cards[index] + 1
                cards[index + 1] = cards[index + 1] + 1
                cards[index + 2] = cards[index + 2] + 1
            end

            ParseOrdinalArray(cards, eye)
            if level < LEVEL then
                GenerateOrdinalArraySub(cards, level + 1, eye)
            end

            if i <= 9 then
                cards[i] = cards[i] - 3
            else
                local index = i - 9
                cards[index] = cards[index] - 1
                cards[index + 1] = cards[index + 1] - 1
                cards[index + 2] = cards[index + 2] - 1
            end
        until(true)
    end
end

function ParseOrdinalArray(cards, eye)
    if not CheckAddOrdinal(cards, 0, eye) then
        return
    end
    ParseOrdinalArraySub(cards, 1, eye)
end

function CheckAddOrdinal(cards, gold, eye)
    local key = 0

    for i = 1, 9 do
        key = key * 10 + cards[i]
    end

    if key == 0 then
        return false
    end

    -- 标记一下已检测过的组合
    if eye then
        if CheckOrdinalTwo[gold][key] then
            return false
        end
        CheckOrdinalTwo[gold][key] = true
    else
        if CheckOrdinalOne[gold][key] then
            return false
        end
        CheckOrdinalOne[gold][key] = true
    end

    for i = 1, 9 do
        if cards[i] > 4 then
            return true
        end
    end

    AddMahjongArray(key, gold, eye, true)
    return true
end

function ParseOrdinalArraySub(cards, gold, eye)
    for i = 1, 9 do
        repeat
            if cards[i] == 0 then
                break
            end

            cards[i] = cards[i] - 1

            if not CheckAddOrdinal(cards, gold, eye) then
                cards[i] = cards[i] + 1
                break
            end
            if gold < MaxGold then
                ParseOrdinalArraySub(cards, gold + 1, eye)
            end
            cards[i] = cards[i] + 1
        until(true)
    end
end

-------------------------------------------------------------------------------

function GenerateFengArray()
    local cards = {0, 0, 0, 0, 0, 0, 0}
    GenerateFengArraySub(cards, 1, false)
end

function GenerateFengArrayWithEye()
    local cards = {0, 0, 0, 0, 0, 0, 0}
    for i = 1, 7 do
        cards[i] = 2
        ParseFengArray(cards, true)
        GenerateFengArraySub(cards, 1, true)
        cards[i] = 0
    end
end

function GenerateFengArraySub(cards, level, eye)
    for i = 1, 7 do
        repeat
            if cards[i] > 3 then
                break
            end
            cards[i] = cards[i] + 3
            ParseFengArray(cards, eye)
            if level < LEVEL then
                GenerateFengArraySub(cards, level + 1, eye)
            end
            cards[i] = cards[i] - 3
        until(true)
    end
end

function ParseFengArray(cards, eye)
    if not CheckAddFeng(cards, 0, eye) then
        return
    end
    ParseFengArraySub(cards, 1, eye)
end

function CheckAddFeng(cards, gold, eye)
    local key = 0

    for i = 1, 7 do
        key = key * 10 + cards[i]
    end

    if key == 0 then
        return false
    end

    -- 标记一下已检测过的组合
    if eye then
        if CheckFengTwo[gold][key] then
            return false
        end
        CheckFengTwo[gold][key] = true
    else
        if CheckFengOne[gold][key] then
            return false
        end
        CheckFengOne[gold][key] = true
    end

    for i = 1, 7 do
        if cards[i] > 4 then
            return true
        end
    end
    AddMahjongArray(key, gold, eye, false)
    return true
end

function ParseFengArraySub(cards, gold, eye)
    for i = 1, 7 do
        repeat
            if cards[i] == 0 then
                break
            end

            cards[i] = cards[i] - 1

            if not CheckAddFeng(cards, gold, eye) then
                cards[i] = cards[i] + 1
                break
            end
            if gold < MaxGold then
                ParseFengArraySub(cards, gold + 1, eye)
            end
            cards[i] = cards[i] + 1
        until(true)
    end
end

-------------------------------------------------------------------------------

-- gold:金牌数 ordinal:是否序数牌 eye:是否带将
function AddMahjongArray(key, gold, eye, ordinal)
    if ordinal then
        if eye then
            MahjongArray.ordinalEyeArray[gold][key] = true
        else
            MahjongArray.ordinalArray[gold][key] = true
        end
    else
        if eye then
            MahjongArray.windEyeArray[gold][key] = true
        else
            MahjongArray.windArray[gold][key] = true
        end
    end
end

-- gold:金牌数 ordinal:是否序数牌 eye:是否带将
function CheckMahjongArray(key, gold, eye, ordinal)
    if ordinal then
        if eye then
            return MahjongArray.ordinalEyeArray[gold][key]
        else
            return MahjongArray.ordinalArray[gold][key]
        end
    else
        if eye then
            return MahjongArray.windEyeArray[gold][key]
        else
            return MahjongArray.windArray[gold][key]
        end
    end
end

-------------------------------------------------------------------------------

--[[
    功能: 检测手牌是否成胡 通用型(3n+2)
    参数:
          classifyCard: 分类组装好的手牌 eg:{[1]={0,0,0,0,0,0,0,0,0}, [2]={0,0,0,0,0,0,0,0,0}, [3]={0,0,0,0,0,0,0,0,0}, [4]={0,0,0,0,0,0,0}}
                        1:万(1~9)  2:条(1~9)  3:筒(1~9)  4:风(1~7)
                        以各类牌的数量替换对应的0
          gold:金牌数量
]]
function CheckWin(classifyCard, gold)
    for i, cards in pairs(classifyCard) do
        local curgold = gold
        local ret, need = CheckNeedGold(cards, curgold, true, i < 4)
        if ret then
            curgold = curgold - need
            for i2, cards2 in pairs(classifyCard) do
                if i2 ~= i then
                    ret, need = CheckNeedGold(cards2, curgold, false, i2 < 4)
                    if not ret then
                        break
                    end
                    curgold = curgold - need
                end
            end
        end
        if ret then
            return true
        end
    end
    return false
end

function CheckNeedGold(cards, gold, eye, ordinal)
    local key, num = 0, 0
    for k, v in ipairs(cards) do
        num = num + v
        key = key * 10 + v
    end

    for n = 0, gold do
        local remainder = (num + n)%3
        if eye then
            if remainder == 2 and CheckMahjongArray(key, n, eye, ordinal) then
                return true, n
            end
        else
            if remainder == 0 and CheckMahjongArray(key, n, eye, ordinal) then
                return true, n
            end
        end
    end

    return false
end
