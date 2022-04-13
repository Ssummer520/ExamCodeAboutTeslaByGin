package conf

import (
	"sync"
)
import "ExamCode/models"

var CarInventoryModel models.CarInventory
var CarLock sync.RWMutex
var Lock sync.RWMutex
var Age = 1
