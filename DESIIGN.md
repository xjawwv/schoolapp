# Design System

## Philosophy

Design decisions are made with intent. Every spacing value, font pairing, and color choice should feel like it was made by a person with a clear point of view — not generated. Avoid defaults. Avoid the middle. Commit to a direction.

---

## Anti-Patterns

These are forbidden. They produce generic, forgettable interfaces.

- `Inter`, `Roboto`, `Arial`, `system-ui` as primary typefaces
- Purple or violet gradients on white backgrounds
- Soft blue as the default accent color
- Rounded cards with drop shadows as the primary layout pattern
- Icon + label + chevron as the default list item structure
- Glass morphism for decorative purposes
- Pastel palettes with no dominant color
- Centered hero with a headline, subtext, and two buttons

---

## Typography

Choose fonts that carry character. Each pairing should feel like a deliberate editorial decision.

### Approved Pairings

| Role          | Font                  | Source              |
|---------------|-----------------------|---------------------|
| Display       | `Instrument Serif`    | Google Fonts        |
| Display       | `DM Serif Display`    | Google Fonts        |
| Display       | `Playfair Display`    | Google Fonts        |
| Body          | `Geist`               | Vercel              |
| Body          | `Lato`                | Google Fonts        |
| Body          | `Source Serif 4`      | Google Fonts        |
| Mono          | `Geist Mono`          | Vercel              |
| Mono          | `Fira Code`           | Google Fonts        |

### Scale

```css
--text-xs:   0.75rem;
--text-sm:   0.875rem;
--text-base: 1rem;
--text-lg:   1.125rem;
--text-xl:   1.25rem;
--text-2xl:  1.5rem;
--text-3xl:  1.875rem;
--text-4xl:  2.25rem;
--text-5xl:  3rem;
```

### Rules

- Display font is used for headings only — never body copy
- Line height for body is `1.6`, for headings `1.1`
- Letter spacing for display text: `-0.02em`
- Letter spacing for UI labels: `0.04em` with `text-transform: uppercase`
- Maximum line length: `68ch`

---

## Color

Color communicates hierarchy and tone. Pick one dominant color and one sharp accent. Neutrals carry the rest.

### Construction

```css
--color-bg:        #0e0e0e;
--color-surface:   #161616;
--color-border:    #2a2a2a;
--color-muted:     #5a5a5a;
--color-text:      #e8e8e8;
--color-heading:   #ffffff;
--color-accent:    #e8c547;
--color-accent-fg: #0e0e0e;
--color-error:     #e05252;
--color-success:   #4caf7d;
```

### Principles

- One accent color per interface. No secondary accent unless the design explicitly requires contrast.
- Dark backgrounds are preferred for data-dense or technical interfaces
- Light interfaces use warm off-whites (`#f5f4f0`, `#faf9f6`), never pure white
- Color must never be the only differentiator for status — pair it with shape or text

---

## Spacing

All spacing derives from a base unit of `4px`.

```css
--space-1:  4px;
--space-2:  8px;
--space-3:  12px;
--space-4:  16px;
--space-5:  20px;
--space-6:  24px;
--space-8:  32px;
--space-10: 40px;
--space-12: 48px;
--space-16: 64px;
--space-20: 80px;
--space-24: 96px;
```

Padding inside components follows strict rules:

| Component Type  | Padding           |
|-----------------|-------------------|
| Button (sm)     | `6px 12px`        |
| Button (md)     | `8px 16px`        |
| Button (lg)     | `12px 24px`       |
| Card            | `24px`            |
| Input           | `8px 12px`        |
| Modal           | `32px`            |
| Section         | `80px 0`          |

---

## Layout

Layouts should be structural and opinionated. Avoid centering everything.

### Grid

- Base column grid: 12 columns
- Column gap: `24px`
- Max content width: `1200px`
- Page padding (mobile): `16px`
- Page padding (tablet): `32px`
- Page padding (desktop): `48px`

### Principles

- Left-aligned text is default. Center-aligned text is reserved for isolation moments (empty states, modal confirmations).
- Asymmetric layouts are encouraged for marketing and editorial surfaces
- Data tables and dashboards use full-width, edge-to-edge layouts
- Sidebar widths: `240px` (collapsed: `64px`)

---

## Components

### Buttons

Buttons have no border radius by default. Use `2px` only for pill variants.

```css
.button-primary {
  background: var(--color-accent);
  color: var(--color-accent-fg);
  font-family: var(--font-body);
  font-size: var(--text-sm);
  font-weight: 600;
  letter-spacing: 0.03em;
  text-transform: uppercase;
  padding: 8px 16px;
  border: none;
  border-radius: 2px;
  cursor: pointer;
}

.button-ghost {
  background: transparent;
  color: var(--color-text);
  border: 1px solid var(--color-border);
  border-radius: 2px;
  padding: 8px 16px;
}
```

### Cards

Cards use a visible left border as a structural anchor — not a shadow.

```css
.card {
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  border-left: 3px solid var(--color-accent);
  padding: var(--space-6);
  border-radius: 0;
}
```

### Inputs

```css
.input {
  background: var(--color-bg);
  border: 1px solid var(--color-border);
  border-radius: 0;
  color: var(--color-text);
  font-family: var(--font-body);
  font-size: var(--text-sm);
  padding: 8px 12px;
  transition: border-color 120ms ease;
}

.input:focus {
  border-color: var(--color-accent);
  outline: none;
}
```

---

## Motion

Animation is used for state transitions, not decoration.

| Interaction         | Duration | Easing              |
|---------------------|----------|---------------------|
| Hover state         | 80ms     | `linear`            |
| Focus ring          | 100ms    | `ease-out`          |
| Dropdown open       | 140ms    | `ease-out`          |
| Modal enter         | 200ms    | `cubic-bezier(0.16, 1, 0.3, 1)` |
| Page transition     | 300ms    | `ease-in-out`       |
| Skeleton pulse      | 1.6s     | `ease-in-out`       |

Never animate layout properties (`width`, `height`, `top`, `left`). Use `transform` and `opacity`.

---

## Iconography

All icons come from `lucide-react`. No other icon libraries are permitted.

- Icons must not be used as decoration without semantic purpose
- Icon size must follow the defined scale (see `agents.md`)
- Icons in buttons are always paired with a visible text label (exception: icon-only toolbar buttons with tooltips)
- No filled icons — use stroke-only variants exclusively

---

## Elevation

Shadows are used sparingly. Prefer borders.

```css
--shadow-sm:  0 1px 2px rgba(0,0,0,0.4);
--shadow-md:  0 4px 12px rgba(0,0,0,0.5);
--shadow-lg:  0 12px 32px rgba(0,0,0,0.6);
```

Use `--shadow-md` for modals and floating panels only. Everything else uses border separation.

---

## Tone Reference

When making design decisions, ask: does this look like it was made by a team with taste, or generated by a model trained on average interfaces?

The answer should always be the former.
