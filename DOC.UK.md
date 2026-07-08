# mistral - довідник

Повний довідник пакета `mistral`: клієнт, спільна модель `goloop/ai`, chat
completions (інтерфейс і нативний), стрімінг, embeddings і моделі.

Англійська версія: **[DOC.md](DOC.md)**.

## Зміст

- [Ментальна модель](#ментальна-модель)
- [Створення клієнта](#створення-клієнта)
- [Generate і Stream](#generate-і-stream)
- [Нативні chat completions](#нативні-chat-completions)
- [Інструменти, зображення й system-промпти](#інструменти-зображення-й-system-промпти)
- [Embeddings](#embeddings)
- [Моделі](#моделі)
- [Опції та помилки](#опції-та-помилки)

## Ментальна модель

`mistral.Client` реалізує `ai.Client` - провайдер-незалежний контракт із
`github.com/goloop/ai`. Спільні `Generate` і `Stream` покривають спільну основу
(чат із інструментами, зображеннями й стрімінгом), тож код проти інтерфейсу
працює з будь-яким провайдером.

Специфіка провайдера - у нативних методах: повний `ChatCompletion`, embeddings і
перелік моделей. Формат обміну - сумісний із chat completions.

```go
import (
	"github.com/goloop/ai"
	"github.com/goloop/mistral"
)
```

## Створення клієнта

```go
c := mistral.New(os.Getenv("MISTRAL_API_KEY"))

c = mistral.New(apiKey, mistral.WithTimeout(30*time.Second))
```

Base URL за замовчуванням `https://api.mistral.ai/v1`. Наведіть `WithBaseURL` на
будь-який сумісний ендпоінт, щоб перевикористати клієнт.

## Generate і Stream

```go
resp, err := c.Generate(ctx, &ai.Request{
	Model:    mistral.ModelLargeLatest,
	System:   "You are concise.",
	Messages: []ai.Message{ai.UserText("Name three primary colors.")},
})
resp.Text()
resp.ToolCalls()
resp.Usage
```

`Stream` повертає `iter.Seq2[ai.Chunk, error]`: текстові дельти чанками з `Text`,
завершений виклик інструмента - чанком із `ToolCall`, фінальний чанк - `Done` і
`Usage`.

## Нативні chat completions

Для опцій, специфічних для провайдера, будуйте `ChatRequest` і викликайте
`ChatCompletion` чи `ChatCompletionStream`:

```go
resp, err := c.ChatCompletion(ctx, &mistral.ChatRequest{
	Model:          mistral.ModelLargeLatest,
	Messages:       []mistral.ChatMessage{{Role: "user", Content: "as JSON"}},
	ResponseFormat: json.RawMessage(`{"type":"json_object"}`),
})
```

## Інструменти, зображення й system-промпти

Інструменти, зображення й system-промпти використовують спільні типи `ai`:
`ai.Tool`, `ai.Image`, `ai.ToolResult` і повідомлення `RoleSystem` або поле
`System`. Результати інструментів надсилаються назад повідомленнями `RoleTool`,
де `ai.ToolResult.ID` збігається з `ai.ToolUse.ID`. Вбудовані байти зображення
надсилаються як base64 data URI (потрібна vision-модель, напр. `ModelPixtral`).

## Embeddings

```go
vecs, err := c.Embed(ctx, mistral.ModelEmbed, "hello", "world")
resp, err := c.Embeddings(ctx, &mistral.EmbeddingRequest{
	Model: mistral.ModelEmbed, Input: []string{"hello"},
})
```

## Моделі

```go
models, err := c.Models(ctx)
m, err := c.GetModel(ctx, mistral.ModelLargeLatest)
```

## Опції та помилки

Опції: `WithBaseURL`, `WithHTTPClient`, `WithTimeout`, `WithMaxRetries`,
`WithHeader`.

Невдала відповідь стає `*ai.APIError` зі `Status`, `Type`, `Code`, `Message` і
сирим тілом:

```go
var apiErr *ai.APIError
if errors.As(err, &apiErr) && apiErr.Status == http.StatusTooManyRequests {
	// backoff
}
```

Запити без моделі чи повідомлень падають до мережі з `ai.ErrNoModel` або
`ai.ErrNoMessages`.
