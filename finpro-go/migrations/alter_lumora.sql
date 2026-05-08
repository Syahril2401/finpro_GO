-- ============================================================
-- LUMORA: ALTER result_summary + seed questions
-- Run this against your pt_lumora database
-- ============================================================

-- 1. Fix result_summary: add per-dimension score columns
ALTER TABLE `result_summary`
    ADD COLUMN `planning_score`        INT NULL AFTER `total_score`,
    ADD COLUMN `time_management_score` INT NULL AFTER `planning_score`,
    ADD COLUMN `cognitive_score`       INT NULL AFTER `time_management_score`,
    ADD COLUMN `reflection_score`      INT NULL AFTER `cognitive_score`;

-- 2. Fix result_summary: category_result must hold long AI JSON (not just 150 chars)
ALTER TABLE `result_summary`
    MODIFY COLUMN `category_result` TEXT;

-- 3. Fix ai_logs: ai_output must accept large JSON blobs
ALTER TABLE `ai_logs`
    MODIFY COLUMN `ai_output` LONGTEXT;

-- 4. Seed 20 SRL questions (no admin FK needed — created_by = NULL)
--    Each question has a category used for per-dimension score calculation.
--    UUIDs are fixed so the frontend can always reference them.

INSERT IGNORE INTO `questions`
    (`question_id`, `question_text`, `category`, `scale_min`, `scale_max`, `created_by`, `is_active`)
VALUES
    -- PLANNING (Q1-Q5)
    ('11111111-0001-0001-0001-000000000001', 'Saya membuat jadwal belajar sebelum mulai belajar',        'planning', 1, 5, NULL, 1),
    ('11111111-0001-0001-0001-000000000002', 'Saya menetapkan target belajar yang jelas',                 'planning', 1, 5, NULL, 1),
    ('11111111-0001-0001-0001-000000000003', 'Saya merencanakan materi yang akan dipelajari',             'planning', 1, 5, NULL, 1),
    ('11111111-0001-0001-0001-000000000004', 'Saya menentukan waktu belajar secara rutin',                'planning', 1, 5, NULL, 1),
    ('11111111-0001-0001-0001-000000000005', 'Saya mempersiapkan kebutuhan belajar sebelum mulai',        'planning', 1, 5, NULL, 1),
    -- TIME MANAGEMENT (Q6-Q10)
    ('11111111-0001-0001-0002-000000000006', 'Saya mengatur waktu belajar dengan baik',                   'time_management', 1, 5, NULL, 1),
    ('11111111-0001-0001-0002-000000000007', 'Saya menyelesaikan tugas tepat waktu',                      'time_management', 1, 5, NULL, 1),
    ('11111111-0001-0001-0002-000000000008', 'Saya jarang menunda pekerjaan',                             'time_management', 1, 5, NULL, 1),
    ('11111111-0001-0001-0002-000000000009', 'Saya memprioritaskan tugas yang penting',                   'time_management', 1, 5, NULL, 1),
    ('11111111-0001-0001-0002-000000000010', 'Saya konsisten dengan jadwal belajar saya',                 'time_management', 1, 5, NULL, 1),
    -- COGNITIVE (Q11-Q15)
    ('11111111-0001-0001-0003-000000000011', 'Saya menggunakan metode belajar tertentu (mencatat, merangkum, dll)', 'cognitive', 1, 5, NULL, 1),
    ('11111111-0001-0001-0003-000000000012', 'Saya mencoba berbagai cara belajar untuk menemukan yang efektif',     'cognitive', 1, 5, NULL, 1),
    ('11111111-0001-0001-0003-000000000013', 'Saya memahami materi, bukan hanya menghafal',                         'cognitive', 1, 5, NULL, 1),
    ('11111111-0001-0001-0003-000000000014', 'Saya mengulang materi untuk memperkuat pemahaman',                    'cognitive', 1, 5, NULL, 1),
    ('11111111-0001-0001-0003-000000000015', 'Saya menghubungkan materi dengan pengetahuan sebelumnya',             'cognitive', 1, 5, NULL, 1),
    -- REFLECTION (Q16-Q20)
    ('11111111-0001-0001-0004-000000000016', 'Saya mengecek apakah saya memahami materi',                           'reflection', 1, 5, NULL, 1),
    ('11111111-0001-0001-0004-000000000017', 'Saya menyadari ketika saya tidak memahami sesuatu',                   'reflection', 1, 5, NULL, 1),
    ('11111111-0001-0001-0004-000000000018', 'Saya mengevaluasi cara belajar saya',                                 'reflection', 1, 5, NULL, 1),
    ('11111111-0001-0001-0004-000000000019', 'Saya memperbaiki strategi belajar jika kurang efektif',               'reflection', 1, 5, NULL, 1),
    ('11111111-0001-0001-0004-000000000020', 'Saya belajar dari kesalahan sebelumnya',                              'reflection', 1, 5, NULL, 1);

SELECT CONCAT('✅ Seeded ', COUNT(*), ' questions') AS status FROM questions;
SELECT CONCAT('✅ result_summary columns: ', GROUP_CONCAT(COLUMN_NAME ORDER BY ORDINAL_POSITION)) AS columns
FROM INFORMATION_SCHEMA.COLUMNS
WHERE TABLE_NAME = 'result_summary' AND TABLE_SCHEMA = DATABASE();
