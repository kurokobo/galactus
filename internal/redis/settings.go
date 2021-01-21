package redis

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/automuteus/utils/pkg/rediskey"
	"github.com/automuteus/utils/pkg/settings"
	"github.com/go-redis/redis/v8"
)

func GetSettingsFromRedis(client *redis.Client, guildID string) (*settings.GuildSettings, error) {
	var sett settings.GuildSettings
	key := rediskey.GuildSettings(hashGuildID(guildID))

	str, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(str), &sett)
	if err != nil {
		return nil, err
	}
	return &sett, nil
}

func hashGuildID(guildID string) string {
	return genericHash(guildID)
}

func genericHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}