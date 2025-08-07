# 🧠 How to Write Proper Git Commit Messages

Writing clean and meaningful commit messages helps your team and future self understand the **why** behind code changes.

---

## ✅ 1. Basic Git Commit Command

```bash
git commit -m "Your concise commit message"
```

---

## ✅ 2. Conventional Commit Format (Recommended)

**Structure:**
```
<type>(optional-scope): <short summary>

<optional detailed description>

<optional footer>
```

**Example:**
```bash
git commit -m "feat(auth): add login with Google"
```

Or with full message:

```bash
git commit
```

Then inside the editor:

```
feat(auth): add login with Google

Added OAuth2 flow using Google Sign-In.
This allows users to authenticate using their Google accounts.

Fixes #42
```

---

## 🔥 Common Commit Types

| Type       | Meaning                                  |
|------------|------------------------------------------|
| `feat`     | A new feature                            |
| `fix`      | A bug fix                                |
| `docs`     | Documentation changes                    |
| `style`    | Formatting (no code changes)             |
| `refactor` | Refactor code (no feature/bug fix)       |
| `test`     | Add or fix tests                         |
| `chore`    | Tooling or build changes                 |
| `perf`     | Performance improvements                 |

---

## ✅ 3. Good Commit Message Examples

- `feat(api): implement CRUD for users`
- `fix(order): resolve total price calculation bug`
- `docs: update README with setup instructions`
- `refactor: clean up nested if-else in controller`

---

## ❌ Avoid These

- `update stuff`
- `fix`
- `changes`
- `wip` (unless you're deliberately pushing work in progress)

---

## Bonus: Git Emoji Style (Optional)

```bash
git commit -m "✨ feat(ui): add dark mode toggle"
```

| Emoji | Meaning         |
|-------|-----------------|
| ✨     | Feature         |
| 🐛     | Bug fix         |
| 📝     | Documentation   |
| ♻️     | Refactor        |
| 🚀     | Deployment      |

---

**Keep your history clean and readable. Your team will thank you.**