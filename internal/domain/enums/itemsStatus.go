package enums

type ItemStatus string

const (
    StatusTersedia  ItemStatus = "Tersedia"
    StatusTerjual   ItemStatus = "Terjual"
    StatusRusak     ItemStatus = "Rusak"
    StatusTerpinjam ItemStatus = "Terpinjam"
    StatusLainnya   ItemStatus = "Lainnya"
)
