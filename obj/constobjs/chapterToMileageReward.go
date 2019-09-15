package constobjs

import (
    "strconv"

    "github.com/fluofoxxo/outrun/enums"
    "github.com/fluofoxxo/outrun/obj"
)

var AreaRewards = map[string][]obj.MileageReward{
    "1,1": []obj.MileageReward{
        invincibleItem(1),
        invincibleItem(3),
    },
}

func GetAreaReward(chapter, episode int64) []obj.MileageReward {
    chapS := strconv.Itoa(int(chapter))
    epS := strconv.Itoa(int(episode))
    getS := chapS + "," + epS
    return AreaRewards[getS]
}

func invincibleItem(point int64) obj.MileageReward {
    return obj.NewMileageReward(enums.ItemIDInvincible, point)
}
