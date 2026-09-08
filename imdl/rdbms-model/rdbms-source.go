package rdbmsmodel

// SourceRDBMS describes the source RDBMS environment/group input for recommendation and model persistence.
type SourceRDBMS struct {
	Description          string                `json:"description,omitempty"`
	SourceCloud          *CloudProperty        `json:"sourceCloud,omitempty"`
	SourceRDBMSInstances []SourceRDBMSProperty `json:"sourceRDBMSInstances" validate:"required,min=1"`
}

// SourceRDBMSProperty describes one RDBMS instance observed in the source environment.
type SourceRDBMSProperty struct {
	// DisplayName is an optional human-friendly alias or cloud DB identifier (e.g. "billing-db").
	DisplayName string `json:"displayName,omitempty" example:"billing-db"`

	// Description is an optional free-form note or business purpose.
	Description string `json:"description,omitempty" example:"Production billing service master database"`

	// DBNode represents the underlying host/server infrastructure specs (CPU, Memory, Disks).
	DBNode DBNodeProperty `json:"dbNode" validate:"required"`

	// DBEngine represents the database engine software (type, version, port, role, HA).
	DBEngine DBEngineProperty `json:"dbEngine" validate:"required"`

	// InnerDatabases lists the logical tenant databases residing inside this RDBMS instance.
	InnerDatabases []InnerDatabaseProperty `json:"innerDatabases,omitempty"`
}

// DBNodeProperty represents the computing server/node infrastructure hosting the database.
type DBNodeProperty struct {
	Hostname  string         `json:"hostname,omitempty" example:"db-node-01"`
	MachineId string         `json:"machineId,omitempty" example:"node-550e8400-e29b-41d4-a716-446655440000"`
	CPU       CpuProperty    `json:"cpu" validate:"required"`
	Memory    MemoryProperty `json:"memory" validate:"required"`
	RootDisk  DiskProperty   `json:"rootDisk,omitempty"`
	DataDisks []DiskProperty `json:"dataDisks,omitempty"`
}

// CpuProperty represents CPU specifications of the database host machine.
type CpuProperty struct {
	Architecture string  `json:"architecture,omitempty" example:"x86_64"`
	Cpus         uint32  `json:"cpus" validate:"required" example:"1"`    // Number of physical CPUs (sockets)
	Cores        uint32  `json:"cores" validate:"required" example:"2"`   // Number of physical cores per CPU
	Threads      uint32  `json:"threads" validate:"required" example:"2"` // Number of logical CPUs (threads) per CPU
	MaxSpeed     float32 `json:"maxSpeed,omitempty" example:"3.6"`        // Maximum speed in GHz
	Vendor       string  `json:"vendor,omitempty" example:"GenuineIntel"`
	Model        string  `json:"model,omitempty" example:"Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz"`
}

// MemoryProperty represents memory specifications of the database host machine.
type MemoryProperty struct {
	Type      string `json:"type,omitempty" example:"DDR4"`
	TotalSize uint64 `json:"totalSize" validate:"required" example:"4"` // Unit GiB
	Available uint64 `json:"available,omitempty"`                       // Unit GiB
	Used      uint64 `json:"used,omitempty"`                            // Unit GiB
}

// DiskProperty represents disk/storage specifications of the database host machine.
type DiskProperty struct {
	Label     string `json:"label,omitempty" example:"/"`
	Type      string `json:"type" validate:"required" example:"SSD"`      // SSD, HDD
	TotalSize uint64 `json:"totalSize" validate:"required" example:"100"` // Unit GB
	Available uint64 `json:"available,omitempty"`                         // Unit GB
	Used      uint64 `json:"used,omitempty"`                              // Unit GB
}

// DBEngineProperty describes the database engine software configuration and role.
type DBEngineProperty struct {
	Engine        string `json:"engine" validate:"required" example:"mysql"`      // "mysql", "mariadb", "postgresql"
	EngineVersion string `json:"engineVersion" validate:"required" example:"8.0"` // e.g. "8.0", "10.5"
	Port          int    `json:"port,omitempty" example:"3306"`                   // DB port (default: 3306)
	Role          string `json:"role,omitempty" example:"primary"`                // "primary" (writer), "replica" (reader), "standalone"
}

// InnerDatabaseProperty describes a logical database inside the RDBMS instance.
type InnerDatabaseProperty struct {
	DatabaseName string  `json:"databaseName" validate:"required" example:"sampledb"`
	CharacterSet string  `json:"characterSet,omitempty" example:"utf8mb4"`
	Collation    string  `json:"collation,omitempty" example:"utf8mb4_unicode_ci"`
	SizeMb       float64 `json:"sizeMb,omitempty" example:"512.5"`
	TableCount   int     `json:"tableCount,omitempty" example:"24"`
	RowCount     int64   `json:"rowCount,omitempty" example:"150000"`
}
