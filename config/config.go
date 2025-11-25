package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

// Config 配置管理结构体
type Config struct {
	Mode   string
	DBPath string
}

// 全局配置实例
var (
	globalConfig *Config
	once         sync.Once
)

// GetConfig 获取配置实例
func GetConfig() *Config {
	once.Do(func() {
		loadConfig()
	})
	return globalConfig
}

// GetMode 获取运行模式 example: release, dev
func GetMode() string {
	return GetConfig().Mode
}

// IsDev 判断是否为开发模式
func IsDev() bool {
	return GetMode() == "dev"
}

// GetDBPath 获取数据库路径 example: database.db
func GetDBPath() string {
	return GetConfig().DBPath
}

// getWithDefault 使用泛型的配置获取函数（内部使用）
func getWithDefault[T any](envMap map[string]string, key string, defaultValue T) T {
	value, exists := envMap[key]
	if !exists {
		fmt.Printf("[WARN] config: key '%s' not set, using default '%v'\n", key, defaultValue)
		return defaultValue
	}

	// 根据默认值的类型进行转换
	switch any(defaultValue).(type) {
	case string:
		return any(value).(T)
	case int:
		if intVal, err := strconv.Atoi(value); err == nil {
			return any(intVal).(T)
		}
		fmt.Printf("[WARN] config: key '%s' value '%s' invalid int, using default '%v'\n", key, value, defaultValue)
		return defaultValue
	case bool:
		lowerValue := strings.ToLower(strings.TrimSpace(value))
		switch lowerValue {
		case "true", "1", "yes", "on":
			return any(true).(T)
		case "false", "0", "no", "off":
			return any(false).(T)
		default:
			fmt.Printf("[WARN] config: key '%s' value '%s' invalid bool, using default '%v'\n", key, value, defaultValue)
			return defaultValue
		}
	case float64:
		if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
			return any(floatVal).(T)
		}
		fmt.Printf("[WARN] config: key '%s' value '%s' invalid float, using default '%v'\n", key, value, defaultValue)
		return defaultValue
	default:
		fmt.Printf("[WARN] config: key '%s' unsupported type, using default '%v'\n", key, defaultValue)
		return defaultValue
	}
}

// loadConfig 加载配置
func loadConfig() {
	// 读取 .env 文件
	envMap := make(map[string]string)

	// 获取当前工作目录
	wd, _ := os.Getwd()
	// 构建 .env 文件路径
	envPath := filepath.Join(wd, ".env")
	// 打开文件
	file, err := os.Open(envPath)
	if err != nil {
		// 如果文件不存在，使用默认配置
		globalConfig = &Config{
			Mode: "dev",
		}
		return
	}
	defer file.Close()

	// 逐行读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 跳过空行和注释行
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// 解析键值对
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			quoted := false
			if len(value) >= 2 {
				if (value[0] == '"' && value[len(value)-1] == '"') ||
					(value[0] == '\'' && value[len(value)-1] == '\'') {
					value = value[1 : len(value)-1]
					quoted = true
				}
			}
			if !quoted {
				if idx := strings.Index(value, "#"); idx >= 0 {
					value = strings.TrimSpace(value[:idx])
				}
			}

			envMap[key] = value
		}
	}

	// 创建配置实例
	globalConfig = &Config{
		Mode:   getWithDefault(envMap, "mode", "dev"),
		DBPath: getWithDefault(envMap, "db_path", "database.db"),
	}
}
