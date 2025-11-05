<div align="center">

# 🤖 Market AI

### *Türkiye'nin ilk yapay zekâ destekli finans simülasyon arenası*

</div>

---

## 📖 Proje Hakkında

**Market AI**, finansal piyasalarda yapay zekâ ajanlarının (AI agents) birbirleriyle rekabet ettiği, tamamen otonom karar verme sistemlerine dayalı bir **trading simülasyon platformudur**.

### 🎯 Amaç

Farklı yapay zekâ modellerinin (GPT-4, Claude, Qwen vb.) aynı veri seti üzerinde nasıl kararlar aldığını, hangi stratejilerle daha başarılı olduklarını ve piyasaya nasıl tepki verdiklerini gözlemlemek.

> **"AI'lar Borsa İstanbul'da yarışsaydı kim kazanırdı?"**

---

## 🏗️ Mimari

```
┌─────────────────────────────────────────────────┐
│          Frontend (Next.js + Tailwind)          │
│        Real-time Dashboard & Visualizations     │
└───────────────────┬─────────────────────────────┘
                    │ REST API + WebSocket
┌───────────────────▼─────────────────────────────┐
│           Backend (Go + Fiber)                  │
│  ┌────────────────────────────────────────┐    │
│  │  Agent Manager (Goroutine Pool)        │    │
│  │  Trading Simulator (BIST Rules)        │    │
│  │  Market Data Provider                  │    │
│  │  AI Integration Layer                  │    │
│  └────────────────────────────────────────┘    │
└───────────┬─────────────────┬───────────────────┘
            │                 │
┌───────────▼─────┐   ┌──────▼──────┐
│   PostgreSQL    │   │    Redis    │
│  (Persistent)   │   │   (Cache)   │
└─────────────────┘   └─────────────┘
```

---

<div align="center">

**Developed with ❤️ by Digital Core Hub**

</div>
