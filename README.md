# Create-repo using user access token

## Steps
1. Clone repo
```bash
git clone git@github.com:pansachin/create-repo.git
cd create-repo
```
2. Update user access token in the scripts.
3. Create a repo in user accout
```bash
go run user/createrepo.go
```

4. Create a repo in org account(Token must have access to do that task)
```bash
go run org/createrepo.go
```

5. Create a repo from a existing template(Concept of upstream in pantheon)
```bash
go run template/createrepo.go
```
