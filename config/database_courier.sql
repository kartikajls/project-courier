
INSERT INTO couriers (user_id,name,email,phone,address,level,status,created_at,updated_at) VALUES
('CR001', 'Budiono Hadi Agung', 'budi.hadi@example.com', '081234560001', 'Jl. Ahmad Yani No. 12, Surabaya', '5', 'active', '2026-01-05 08:00:00', '2026-01-05 08:00:00'),
('CR002', 'Andi Pratama', 'andi.pratama@example.com', '081234560002', 'Jl. Diponegoro No. 15, Sidoarjo', '1', 'active', '2026-01-06 09:15:00', '2026-01-06 09:15:00'),
('CR003', 'Rizky Ramadhan', 'rizky.ramadhan@example.com', '081234560003', 'Jl. Gajah Mada No. 20, Malang', '2', 'active', '2026-01-07 10:30:00', '2026-01-07 10:30:00'),
('CR004', 'Dedi Kurniawan', 'dedi.kurniawan@example.com', '081234560004', 'Jl. Veteran No. 8, Gresik', '3', 'active', '2026-01-08 08:45:00', '2026-01-08 08:45:00'),
('CR005', 'Agus Setiawan', 'agus.setiawan@example.com', '081234560005', 'Jl. Pahlawan No. 32, Pasuruan', '1', 'inactive', '2026-01-09 11:00:00', '2026-02-10 09:00:00'),
('CR006', 'Fajar Nugroho', 'fajar.nugroho@example.com', '081234560006', 'Jl. Pemuda No. 17, Mojokerto', '2', 'active', '2026-01-10 07:30:00', '2026-01-10 07:30:00'),
('CR007', 'Hendra Wijaya', 'hendra.wijaya@example.com', '081234560007', 'Jl. Basuki Rahmat No. 25, Jember', '1', 'active', '2026-01-11 09:00:00', '2026-01-11 09:00:00'),
('CR008', 'Dimas Saputra', 'dimas.saputra@example.com', '081234560008', 'Jl. Soekarno Hatta No. 10, Kediri', '2', 'inactive', '2026-01-12 08:20:00', '2026-03-01 10:15:00'),
('CR009', 'Eko Purnomo', 'eko.purnomo@example.com', '081234560009', 'Jl. Kartini No. 45, Banyuwangi', '1', 'active', '2026-01-13 10:00:00', '2026-01-13 10:00:00'),
('CR010', 'Rudi Hartono', 'rudi.hartono@example.com', '081234560010', 'Jl. Merdeka No. 18, Madiun', '3', 'active', '2026-01-14 08:10:00', '2026-01-14 08:10:00'),
('CR011', 'Arif Hidayat', 'arif.hidayat@example.com', '081234560011', 'Jl. Panglima Sudirman No. 22, Sidoarjo', '4', 'active', '2026-01-15 09:30:00', '2026-01-15 09:30:00');

INSERT INTO levels (level_id, name, description) VALUES
('1', 'Petugas Sortir Barang', 'Memilah, mengelompokkan, dan menyiapkan paket berdasarkan tujuan pengiriman'),
('2', 'Kurir Sepeda Motor', 'Kurir operasional perusahaan dengan menggunakan kendaraan sepeda motor'),
('3', 'Kurir Mobil', 'Kurir operasional perusahaan dengan menggunakan kendaraan mobil'),
('4', 'Koordinator', 'Koordinator operasional kurir'),
('5', 'Supervisor', 'Mengawasi dan mengoordinasikan operasional kurir');

INSERT INTO status (type) VALUES
('Active'),
('Inactive');