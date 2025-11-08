# 🤖 Market AI

### _Türkiye'nin ilk yapay zekâ destekli finans simülasyon arenası_

> **"AI'lar Borsa İstanbul'da yarışsaydı kim kazanırdı?"**

---

## 📖 Proje Hakkında

Market AI, finansal piyasalarda yapay zekâ ajanlarının (AI agents) farklı stratejilerle nasıl kararlar aldığını gözlemlemeyi amaçlayan, deneysel bir simülasyon ve test projesidir.

## 🎯 v0.3 - Haber Entegrasyonlu Otonom AI Ajan Sistemi

### ✨ Yeni Özellikler

- **Otonom AI Ajanları**: 30-60 saniye aralıklarla haber bağlamında kendi kendine karar veren AI ajanları
- **Haber Entegrasyonu**: News API + RSS feeds ile Türkiye finans haberlerinin gerçek zamanlı toplanması
- **AI Model Desteği**: OpenAI (GPT-3.5/GPT-4) ve Anthropic (Claude 3 Haiku/Opus)
- **Risk Yönetimi**: Trade'leri gerçekleştirmeden önce otomatik risk doğrulaması
- **Gerçek Zamanlı Akıl Yürütme Beslemesi**: AI ajanlarının düşünce sürecini canlı izleme
- **Pazar Analiz Paneli**: Son haberleri ve etki seviyelerini gösteren dashboard
- **Veritabanı Desteği**: PostgreSQL'de karar zincirlerinin ve düşünce adımlarının depolanması
- **Redis Cache**: Haber cache'leme (30 dakika TTL) ve hızlı erişim

### 🔄 Sistem Mimarisi

#### Backend (Go)

- **News Aggregator**: 30 dakika aralıklarla yeni haberleri getir → Redis cache → WebSocket broadcast
- **Agent Engine**: Her agent için 30-60s aralıklarla:
  1. Piyasa verisi + haberleri topla
  2. AI client'a isteği gönder (haber bağlamıyla)
  3. Kararı kaydet ve düşünme adımlarını depola
  4. Risk Manager'dan geçir
  5. Trade'i çalıştır / reddet
  6. WebSocket'ten broadcast et
- **Risk Manager**: Confidence > 70%, position < 5%, portfolio risk < 20%
- **AI Clients**: OpenAI + Anthropic entegrasyonu
- **News System**: NewsAPI.org + RSS parser (Bloomberg HT, Investing.com, Dünya)

#### Frontend (Next.js)

- **ReasoningFeed**: Gerçek zamanlı AI karar akışı (güven, risk seviyesi, düşünme adımları)
- **LatestNews**: Piyasa haberleri gösterimi (etki seviyesi, ilgili hisseler, duygu)
- **Dashboard**: Ajanların performansı, P&L takibi, canlı durum

### 📊 Karar Döngüsü

```
News Aggregator (30 dk döngü)
    ↓
    [Getir + Önbellekle]
    ↓
Agent Engine (ajan başına 30-60s rastgele)
    ↓ (her döngüde)
    ├─ Piyasa verisi + son haberleri topla
    ├─ AI'ı bağlamla çağır
    ├─ Kararı + düşünme adımlarını kaydet
    ├─ Risk Manager'dan doğrula
    ├─ Trade'i çalıştır/reddet
    └─ WebSocket ile yayınla
    ↓
Frontend ReasoningFeed + News Panel
    ↓
    [Gerçek zamanlı güncellemeler]
```

### 💰 Maliyet Tahminleri

**Test Modelleri (v0.3 varsayılan):**

- GPT-3.5-turbo: $0.001/istek → ~$2-3/gün
- Claude 3 Haiku: $0.00025/istek → ~$0.5/gün
- **Toplam**: ~$3-5/gün

**Production (opsiyonel):**

- GPT-4-turbo: $0.01/istek → ~$20-30/gün
- Claude 3 Opus: $0.015/istek → ~$15-25/gün
- **Toplam**: ~$35-50/gün

### 🚀 Başlangıç

```bash
# Backend (Go 1.23+)
cd cmd/server
go run main.go

# Frontend (Next.js 16)
cd frontend
npm run dev

# Docker (PostgreSQL + Redis)
docker-compose up -d
```

### 📁 Proje Yapısı

```
market-ai/
├── backend (Go)
│   ├── internal/
│   │   ├── models/        # Veri modelleri
│   │   ├── services/      # İş mantığı
│   │   ├── ai/            # AI istemcileri + promptlama
│   │   ├── news/          # Haber toplama
│   │   ├── cache/         # Redis önbellekleme
│   │   └── config/        # Yapılandırma
│   ├── migrations/        # Veritabanı şemaları
│   └── cmd/server/        # Giriş noktası
├── frontend (Next.js)
│   ├── components/        # React bileşenleri
│   ├── lib/               # Yardımcılar
│   └── app/               # Sayfalar
└── docker-compose.yml     # Servisler
```

### 🔧 Amaç

- Farklı AI modellerini aynı veri/koşullarda karşılaştırmak
- Stratejilerin performansını ve karar alma dinamiklerini analiz etmek
- Backend altyapısını (API, DB, Cache) doğrulamak ve ölçümlemek
- Haber bağlamında yapılan kararların etkisini gözlemlemek

## ⚠️ Uyarı

Bu proje yalnızca deneysel ve eğitim/test amaçlıdır. Buradaki hiçbir çıktı, sinyal veya metrik yatırım tavsiyesi değildir; finansal kararlar için kullanılmamalıdır.

---

## 🚀 v0.4 – Çoklu AI Arena & Skor Tablosu

v0.4 ile sistem tekil ajanlardan rekabetçi çoklu yapay zekâ (8 farklı model) arenasına genişletildi.

### ✅ Hedefler

- 8 AI ajanı (OpenAI GPT-4 / GPT-4o-mini, Claude, Gemini, DeepSeek, Llama Groq, Mixtral, Grok)
- Canlı liderlik tablosu (ROI, Kazanma Oranı, P/L, Toplam Değer)
- Periyodik sıralama hesaplama (ağırlıklı skor formülü)
- WebSocket ile anlık güncelleme yayınları
- İstatistik tabloları: günlük, anlık görüntü, karşılıklı maç (temel şema)

### 🗄 Yeni Veritabanı Tabloları (Migration 005)

- `agent_performance_snapshots` – Saatlik/isteğe bağlı anlık görüntü kayıtları
- `leaderboard_rankings` – Hesaplanmış sıralama ve rozetler
- `agent_matchups` – İki ajan arası kazanma-kaybetme takibi
- `agent_daily_stats` – Günlük toplu metrikler (kazanç, kayıp, hacim, en iyi/kötü işlem)
- Fonksiyon: `update_leaderboard_rankings()` – ROI, Kazanma Oranı, P/L ağırlıklı skor

### 🔢 Sıralama Formülü (Genel Sıralama)

$$ overall = (roi \times 0.4) + (win\_rate \times 0.3) + ((total\_profit\_loss / 1000) \times 0.3) $$

### 🔌 Backend Ekleri

- Yeni AI client dosyaları: `google.go`, `deepseek.go`, `groq.go`, `mistral.go`, `xai.go`
- Skor tablosu servisi: periyodik (env ile ayarlanabilir) güncelleme + WebSocket yayını
- REST endpoint: `GET /api/v1/leaderboard`
- Yapılandırma: `LEADERBOARD_UPDATE_INTERVAL` (saniye)

### 🖥 Frontend Ekleri

- `Leaderboard.tsx` – Canlı tablo, ROI rozetleri, P/L, Kazanma Oranı
- Dashboard entegrasyonu

### 🔑 Ortam Değişkenleri (v0.4)

`.env`:

```env
OPENAI_API_KEY=
ANTHROPIC_API_KEY=
GOOGLE_API_KEY=
DEEPSEEK_API_KEY=
GROQ_API_KEY=
MISTRAL_API_KEY=
XAI_API_KEY=

AI_MODEL_GPT=gpt-4-turbo
AI_MODEL_GPT4_MINI=gpt-4o-mini
AI_MODEL_CLAUDE=claude-3-5-sonnet-20241022
AI_MODEL_GEMINI=gemini-1.5-pro
AI_MODEL_DEEPSEEK=deepseek-chat
AI_MODEL_LLAMA=llama-3.1-70b-versatile
AI_MODEL_MIXTRAL=open-mixtral-8x22b
AI_MODEL_GROK=grok-2-latest

AI_TEMPERATURE=0.7
AI_MAX_TOKENS=1500
LEADERBOARD_UPDATE_INTERVAL=60
```

### 📦 Migration Uygulama

```bash
psql -U marketai -d marketai_dev -f migrations/005_agent_stats.sql
```

### 🌱 Seed – Yeni Ajanlar

```sql
INSERT INTO agents (name, model, status, initial_balance, current_balance) VALUES
('Gemini Pro','gemini-1.5-pro','active',100000,100000),
('DeepSeek V3','deepseek-chat','active',100000,100000),
('GPT-4o Mini','gpt-4o-mini','active',100000,100000),
('Llama 3.1 70B','llama-3.1-70b-versatile','active',100000,100000),
('Mixtral 8x22B','open-mixtral-8x22b','active',100000,100000),
('Grok 2','grok-2-latest','active',100000,100000);

INSERT INTO agent_metrics (agent_id)
SELECT id FROM agents WHERE name IN ('Gemini Pro','DeepSeek V3','GPT-4o Mini','Llama 3.1 70B','Mixtral 8x22B','Grok 2')
ON CONFLICT (agent_id) DO NOTHING;
```

### 🔁 Servis Döngüsü

1. Skor tablosu servisi her interval sonunda `update_leaderboard_rankings()` fonksiyonunu çağırır.
2. Sıralama sonuçlarını WebSocket ile `leaderboard_updated` olarak yayınlar.
3. Frontend `Leaderboard.tsx` ilk veriyi REST'ten çeker, sonra anlık güncellemeleri websocket'ten işler.

### 🧪 Doğrulama

```bash
# REST kontrol
curl http://localhost:8080/api/v1/leaderboard | jq

# WebSocket (örnek wscat)
wscat -c ws://localhost:8080/ws
# Mesaj tipini dinle: leaderboard_updated
```

### 💰 Maliyet Analizi (8 Ajan Tam Güç)

| Model                | Tahmini Maliyet / Gün |
| -------------------- | --------------------- |
| GPT-4 Turbo          | ~$14.40               |
| Claude 3.5 Sonnet    | ~$4.32                |
| Gemini 1.5 Pro       | ~$1.80                |
| Grok-2               | ~$2.88                |
| GPT-4o Mini          | ~$0.22                |
| DeepSeek V3          | ~$0.39                |
| Mixtral 8x22B        | ~$2.88                |
| Llama 3.1 70B (Groq) | $0.00                 |

**Toplam (Full Premium)** ≈ $27/gün (~$810/ay)  
**Minimum (Bütçe Seti)** ≈ $2–5/gün

### 💡 Aşamalı Maliyet Stratejisi

| Faz            | Modeller                              | Günlük Maliyet | Amaç                  |
| -------------- | ------------------------------------- | -------------- | --------------------- |
| Faz 1 (Test)   | GPT-4o Mini, DeepSeek, Mixtral, Llama | ~$2            | Fonksiyonel doğrulama |
| Faz 2 (Demo)   | + Gemini, Claude Haiku                | ~$8            | Demo sunumu           |
| Faz 3 (Üretim) | + GPT-4, Claude Sonnet/Opus, Grok     | ~$27           | Rekabetçi analiz      |

### 🎚 Ortam Bayrakları ile Maliyet Kontrolü

`BUDGET_MODE` ve `ENABLE_PREMIUM_MODELS` bayrakları ile çağrı frekansı ve kayıtlı modelleri yönetebilirsin.

| Değişken                | Varsayılan | Etki                                                                                                            |
| ----------------------- | ---------- | --------------------------------------------------------------------------------------------------------------- |
| `BUDGET_MODE`           | `false`    | `true` ise karar döngüsü 30–60 sn yerine 60–120 sn çalışır (istek sayısı azalır).                              |
| `ENABLE_PREMIUM_MODELS` | `true`     | `false` ise GPT-4, Claude (Sonnet/Opus), Grok kayıt edilmez; yalnızca bütçe dostu modeller aktif kalır.        |

Örnek bütçe ayarı:

```env
BUDGET_MODE=true
ENABLE_PREMIUM_MODELS=false
```

### 🔧 Diğer Tasarruf Teknikleri

- Token azaltımı: `AI_TEMPERATURE` sabit tutup prompt içeriğini minimalize et.
- Anlık görüntü seyrekliği: Snapshot kayıtlarını 1 dk yerine 5 dk yap.
- Dinamik hız: Volatilite düşükken interval uzat, yükselince kısalt.
- Fallback: Premium yanıt hatasında otomatik Mixtral/Llama fallback.

---

## 🎉 v0.5 – Çoklu Kaynak Veri Füzyonu & Güvenilirlik Skorlaması

v0.5 ile **çoklu kaynaklı piyasa verisi toplama**, **duygu analizi**, **güvenilirlik skorlaması** ve **gözlemlenebilirlik metrikleri** eklenerek bağlam-farkındalı AI ticaret ajanları güçlendirildi.

### ✨ Eklenen Özellikler

#### 1. Çoklu Kaynak Veri Toplama

- **Yahoo Finance API**: 15 dakika gecikmeli BIST kotasyonları (toplu çekme)
- **Bloomberg HT Scraper**: Colly ile Türk finans haberleri
- **Twitter API**: BIST sembollerinden bahseden son tweetler (arama)

#### 2. Duygu Analizi

- **OpenAI Destekli Tweet Analizi**: Duygu sınıflandırması (pozitif/negatif/nötr) + güven skoru
- **Hisse Bazlı Toplama**: Ortalama duygu, pozitif/negatif sayıları, en etkili tweet
- **Veritabanı Fonksiyonu**: Zaman pencereli toplamalar için `update_sentiment_aggregate()`

#### 3. Füzyon Servisi

**Temel Yetenekler**:

- Paralel olarak fiyat + haber + tweet çekimi
- Tüm tweetler için duygu analizi
- Güven skorlu fiyat anlık görüntüleri kaydetme
- Tekrarlı API çağrılarını azaltmak için 30 saniyelik önbellek
- `price_sources` ve `twitter_sentiment` tablolarına veri kaydetme

#### 4. Güvenilirlik Skorlaması

**Algoritma**:

```
confidence = clamp(50 + 40*successRate - responsePenalty - variancePenalty, 5, 99.9)
```

- `successRate`: kaynak başına geçmiş çekme başarı oranı
- `responsePenalty`: yavaş yanıtları cezalandırır (>1500ms temel değer)
- `variancePenalty`: gelecekte çapraz kaynak fiyat farklılığı tespiti

**Takip**:

- Bellek içi istatistikler: kaynak başına toplam, başarı sayısı, ortalama süre
- DB otomatik güncelleme: `data_sources` tablosu metrikleri izler (asenkron yazma)

#### 5. Gözlemlenebilirlik & Metrikler

**Endpoint**: `GET /api/v1/metrics`

Tüm veri kaynakları için canlı güvenilirlik metriklerini döner:

```json
{
  "success": true,
  "data": {
    "data_sources": [
      {
        "source_type": "yahoo",
        "source_name": "Yahoo Finance API",
        "is_active": true,
        "total_fetches": 120,
        "success_count": 118,
        "error_count": 2,
        "avg_response_time_ms": 850,
        "status": "active",
        "last_fetch_at": "2025-11-08T12:34:56Z"
      }
    ]
  }
}
```

#### 6. AI Prompt Geliştirmeleri

- `DecisionRequest` piyasa bağlamı alanlarıyla genişletildi: `MCPrices`, `MCSentiments`, `MCTopTweets`, `MCNotes`
- `BuildDecisionPrompt()` canlı fiyatlar, duygu özeti ve en etkili tweetlerle "MARKET CONTEXT" bölümü ekler
- Ajan kararları artık gerçek zamanlı çoklu kaynak verisinden faydalanır

#### 7. Debug Endpoint'leri

**Rotalar**: `/api/v1/debug/{yahoo,scraper,tweets}`

- `/debug/yahoo?symbols=THYAO,AKBNK`: Doğrudan Yahoo fiyat çekimi
- `/debug/scraper`: Bloomberg HT haber kazıması
- `/debug/tweets?max=50&analyze=true`: Opsiyonel duygu analiziyle son tweetler

#### 8. Veritabanı Şeması

**Migration 006** (`migrations/006_data_sources.sql`):

- `data_sources`: Güvenilirlik metrik takibi
- `price_sources`: Çoklu kaynak fiyat anlık görüntüleri + güven
- `twitter_sentiment`: Duygu skorlarıyla tweet arşivi
- `stock_sentiment_aggregates`: Zaman pencereli duygu toplamaları
- `scraped_articles`: Hisse bahisli haber arşivi

**Migration 007** (`migrations/007_data_sources_seed.sql`):

- Metrik takibi için 3 temel veri kaynağı başlatması

#### 9. Servisler & Otomasyon

- **MarketDataCollector**: Periyodik veri çekme (yapılandırılabilir aralıklar)
- **SentimentTracker**: Her N dakikada `update_sentiment_aggregate()` çalıştırır
- **AgentEngine**: Karar döngüsüne `MarketContext` enjekte eder

#### 10. Frontend Bileşenleri

- **MarketDataSources**: Kaynak başına canlı çekme süreleri
- **StockSentimentPanel**: Gerçek zamanlı duygu göstergeleri + en etkili tweet
- **BreakingNews**: Etki seviyeleriyle akan haber beslemesi
- **SentimentGauge**: Renkli duygu göstergesi

### 📊 Sistem Akışı (v0.5)

```
┌─────────────────┐
│  Veri Kaynakları│
│  (Yahoo,        │
│   Bloomberg HT, │
│   Twitter API)  │
└────────┬────────┘
         │ çekme (zamanlanmış)
         ▼
┌─────────────────┐
│ Füzyon Servisi  │◄─── 30s önbellek
│ (MarketContext) │
└────────┬────────┘
         │ kaydetme
         ▼
┌─────────────────┐      ┌─────────────────┐
│  price_sources  │      │twitter_sentiment│
│  (güven)        │      │  (analiz edilmiş)│
└─────────────────┘      └─────────────────┘
         │                        │
         └───────┬────────────────┘
                 │ topla
                 ▼
         ┌───────────────┐
         │stock_sentiment│
         │  _aggregates  │
         └───────┬───────┘
                 │
                 ▼
         ┌───────────────┐
         │  AI Ajanları  │◄─── Gelişmiş Prompt
         │  (Karar       │     MarketContext ile
         │   Motoru)     │
         └───────────────┘
```

### 🚀 v0.5 Kurulum

#### 1. Migration'ları Uygula

```bash
docker exec -i marketai-postgres psql -U marketai -d marketai_dev < migrations/006_data_sources.sql
docker exec -i marketai-postgres psql -U marketai -d marketai_dev < migrations/007_data_sources_seed.sql
```

#### 2. Ortamı Yapılandır

`.env` dosyasına ekle:

```env
# Sembol Evreni
SYMBOL_UNIVERSE=THYAO,AKBNK,ASELS,GARAN,BIMAS,KCHOL,SISE

# Çekme Aralıkları (dakika)
YAHOO_FETCH_INTERVAL=5
SCRAPER_FETCH_INTERVAL=15
TWITTER_FETCH_INTERVAL=10
SENTIMENT_UPDATE_INTERVAL=5

# Twitter API
TWITTER_API_KEY=api_anahtarin
TWITTER_API_SECRET=api_sifren
TWITTER_ACCESS_TOKEN=erisim_tokenin
TWITTER_ACCESS_SECRET=erisim_sifresi
```

#### 3. Derle & Çalıştır

```bash
go build -o bin/market-ai ./cmd/server
./bin/market-ai
```

#### 4. Endpoint'leri Doğrula

```bash
# Piyasa bağlamı
curl "http://localhost:8080/api/v1/market/context?symbols=THYAO,AKBNK" | jq

# Metrikler
curl http://localhost:8080/api/v1/metrics | jq

# Debug
curl "http://localhost:8080/api/v1/debug/yahoo?symbols=THYAO" | jq
curl "http://localhost:8080/api/v1/debug/tweets?max=20&analyze=true" | jq
```

### 🧪 Testler

```bash
# Tüm testleri çalıştır
go test ./...

# Güvenilirlik skorlaması testi
go test ./internal/datasources/fusion -v -run TestComputeConfidence

# Handler testleri
go test ./internal/api/handlers -v
```

### 📚 Dokümantasyon

- **Güvenilirlik Skorlaması**: [docs/GUVENILIRLIK.md](docs/GUVENILIRLIK.md)

### 🎯 v0.5 Başarılar

✅ 3 veri kaynağı (Yahoo, Bloomberg HT, Twitter)  
✅ Duygu analizi (OpenAI destekli + toplama)  
✅ Güvenilirlik skorlaması (fiyat başına 0-100 güven)  
✅ Gözlemlenebilirlik (`/api/v1/metrics` endpoint + DB takibi)  
✅ AI prompt geliştirmesi (bağlam-farkındalı kararlar)  
✅ Debug endpoint'leri (kaynak başına tanılama)  
✅ Otomatik toplayıcılar (zamanlanmış çekimler + duygu güncellemeleri)  
✅ Kapsamlı testler (güvenilirlik skorlaması + handler'lar)  
✅ Üretime hazır build

### 🔮 v0.6 Yol Haritası

- Çoklu kaynak fiyat füzyonu (güvene göre ağırlıklı ortalama)
- Çakışma tespiti ve uyarılar (kaynaklar %5'ten fazla farklıysa)
- Grafana dashboardları için Prometheus exporter
- Dinamik kaynak kapatma (güvenilir olmayan kaynakları otomatik devre dışı bırak)
- Twitter streaming API entegrasyonu (gerçek zamanlı duygu)
- Duygu-odaklı risk ayarlamaları
- Ek scraper'lar (Investing.com, KAP)

---

### 🛡 Notlar

- Ajanlar gerçek para veya gerçek zamanlı canlı piyasa yerine simüle edilmiş veride karar verir.
- Maliyet hesapları tahmini (token/istek hacmine bağlı değişir). Gerçek kullanımda bütçe limiti koyun.

---
