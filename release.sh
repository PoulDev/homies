
lastTag=$(git tag --sort=-v:refname | head -n 1)
echo "Enter tag, last tag: $lastTag"
read tag

echo "Enter commit message:"
read commitmsg

git add .
git tag $tag
git commit -m "$commitmsg"
git push --tags

