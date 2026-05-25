# Agents

## Overview

Agents are AI-powered modules that perform tasks autonomously using tools, memory, and decision logic. Each agent has a defined role, a set of allowed tools, and a scope of responsibility.

---

## Rules

- No inline comments in code
- No emoji anywhere in the interface or output
- Icons must come exclusively from the `lucide-react` library
- Each agent must have a single, well-defined responsibility
- Agents must not share mutable state directly

---

## Icon Usage

All icons are imported from `lucide-react`. Use the following pattern:

```tsx
import { Bot, Cpu, Search, FileText, Settings, RefreshCw, AlertCircle } from "lucide-react"
```

Icon sizing follows a strict scale:

| Context        | Size  | Class          |
|----------------|-------|----------------|
| Inline text    | 14px  | `size-3.5`     |
| Button         | 16px  | `size-4`       |
| Card header    | 20px  | `size-5`       |
| Section header | 24px  | `size-6`       |
| Hero / empty   | 40px  | `size-10`      |

---

## Agent Structure

Each agent is defined with the following shape:

```ts
type Agent = {
  id: string
  name: string
  icon: LucideIcon
  role: string
  tools: string[]
  status: "active" | "idle" | "error"
}
```

---

## Registered Agents

### Research Agent
- **Icon**: `Search` from lucide-react
- **Role**: Retrieves and synthesizes information from external sources
- **Tools**: `web_search`, `web_fetch`, `summarize`
- **Status**: Active when a query is in progress

### Document Agent
- **Icon**: `FileText` from lucide-react
- **Role**: Reads, writes, and transforms document files
- **Tools**: `read_file`, `write_file`, `convert_format`
- **Status**: Active when processing a file operation

### Orchestrator Agent
- **Icon**: `Cpu` from lucide-react
- **Role**: Coordinates task delegation across other agents
- **Tools**: `spawn_agent`, `assign_task`, `collect_results`
- **Status**: Always active when a workflow is running

### Settings Agent
- **Icon**: `Settings` from lucide-react
- **Role**: Manages configuration and user preferences
- **Tools**: `read_config`, `write_config`, `validate_schema`
- **Status**: Idle by default, active on config change

---

## Status Indicators

Use lucide icons to represent agent status in the UI:

| Status  | Icon           | Color token        |
|---------|----------------|--------------------|
| Active  | `RefreshCw`    | `text-green-500`   |
| Idle    | `Bot`          | `text-zinc-400`    |
| Error   | `AlertCircle`  | `text-red-500`     |

---

## Adding a New Agent

1. Define the agent object in `agents/registry.ts`
2. Assign a `LucideIcon` to the `icon` field
3. Register its tools in `tools/registry.ts`
4. Add its status handler in `agents/status.ts`
5. Render it using the `<AgentCard />` component

---

## AgentCard Component

```tsx
import { Bot } from "lucide-react"

type Props = {
  agent: Agent
}

export function AgentCard({ agent }: Props) {
  const Icon = agent.icon

  return (
    <div className="agent-card">
      <Icon size={20} />
      <span>{agent.name}</span>
      <span>{agent.status}</span>
    </div>
  )
}
```

---

## Constraints

- Do not import icons from any source other than `lucide-react`
- Do not add comments to agent logic files
- Do not use emoji as status indicators or labels
- Do not create agents with overlapping tool access unless explicitly designed for redundancy
