# How to Run aws-eye CLI

Quick guide to configure your AWS credentials and run the project.

---

## 🚀 Quick Start (3 Steps)

### Step 1: Set Your AWS Credentials

You have two options. Choose **Option A** for quick testing, or **Option B** for permanent setup.

#### Option A: Environment Variables (Quick Test)

Open your terminal and run:

```bash
export AWS_ACCESS_KEY_ID="your-access-key-here"
export AWS_SECRET_ACCESS_KEY="your-secret-access-key-here"
```

**Example:**
```bash
export AWS_ACCESS_KEY_ID="AKIAIOSFODNN7EXAMPLE"
export AWS_SECRET_ACCESS_KEY="wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
```

**Note:** These will only last for your current terminal session. If you close the terminal, you'll need to set them again.

---

#### Option B: AWS Credentials File (Permanent - Recommended)

This method saves your credentials permanently on your computer.

1. **Create the AWS directory** (if it doesn't exist):
   ```bash
   mkdir -p ~/.aws
   ```

2. **Create/edit the credentials file**:
   ```bash
   nano ~/.aws/credentials
   ```
   
   Or use any text editor. Add this content:
   ```ini
   [default]
   aws_access_key_id = your-access-key-here
   aws_secret_access_key = your-secret-access-key-here
   ```

   **Example:**
   ```ini
   [default]
   aws_access_key_id = AKIAIOSFODNN7EXAMPLE
   aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
   ```

3. **Save the file:**
   - If using `nano`: Press `Ctrl+X`, then `Y`, then `Enter`
   - If using `vim`: Press `Esc`, type `:wq`, then `Enter`
   - If using VS Code or other editor: Just save the file

4. **Set proper permissions** (important for security):
   ```bash
   chmod 600 ~/.aws/credentials
   ```

---

### Step 2: Build the Project

Make sure you're in the project directory:

```bash
cd /Users/51xneeraj/Desktop/100x100K/projects/aws-eye
```

Build the project:

```bash
go build -o aws-eye .
```

This creates an executable file named `aws-eye`.

---

### Step 3: Run the CLI

**Run with default region (ap-south-1):**
```bash
./aws-eye
```

**Run with a different region:**
```bash
./aws-eye --region us-east-1
```

**Run with a different region (example):**
```bash
./aws-eye --region eu-west-1
```

---

## 📋 Complete Example

Here's a complete example from start to finish:

```bash
# 1. Navigate to project directory
cd /Users/51xneeraj/Desktop/100x100K/projects/aws-eye

# 2. Set credentials (Option A - temporary)
export AWS_ACCESS_KEY_ID="AKIAIOSFODNN7EXAMPLE"
export AWS_SECRET_ACCESS_KEY="wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"

# 3. Build (only needed once, or after code changes)
go build -o aws-eye .

# 4. Run
./aws-eye
```

---

## ✅ Expected Output

### If you have EC2 instances:
```
Fetching EC2 instances from region: ap-south-1

INSTANCE ID           STATE        PUBLIC IP        AVAILABILITY ZONE
i-0abc12efg34         running      3.109.21.19      ap-south-1a
i-04bb99281de         stopped      —                ap-south-1b
```

### If you have no instances:
```
Fetching EC2 instances from region: ap-south-1

No instances found.
```

### If credentials are wrong:
```
Error: failed to load AWS config: ...
```

### If permissions are missing:
```
Error: failed to describe instances: ... UnauthorizedOperation ...
```
(If you see this, follow the QUICK_FIX.md guide)

---

## 🔍 Verify Your Credentials Are Set

**Check environment variables:**
```bash
echo $AWS_ACCESS_KEY_ID
echo $AWS_SECRET_ACCESS_KEY
```

**Check credentials file:**
```bash
cat ~/.aws/credentials
```

---

## 🛠️ Troubleshooting

### "command not found: aws-eye"
- Make sure you're in the project directory
- Make sure you ran `go build -o aws-eye .` first
- Try: `./aws-eye` (with the `./` prefix)

### "failed to load AWS config"
- Check that your credentials are set correctly
- Verify the access key and secret key are correct
- Make sure there are no extra spaces or quotes in your credentials

### "UnauthorizedOperation"
- Your IAM user needs the `AmazonEC2ReadOnlyAccess` policy
- See `QUICK_FIX.md` for instructions

### "No instances found"
- This is normal! It means you don't have any EC2 instances in that region
- Try a different region: `./aws-eye --region us-east-1`

---

## 💡 Tips

1. **Use Option B (credentials file)** if you plan to use this tool regularly
2. **Never share your credentials** - they're like passwords
3. **The username** you mentioned isn't needed for the CLI - only the access key and secret key are used
4. **Different regions** may have different instances - try multiple regions if needed

---

## 🎯 Quick Reference

```bash
# Set credentials (temporary)
export AWS_ACCESS_KEY_ID="your-key"
export AWS_SECRET_ACCESS_KEY="your-secret"

# Build
go build -o aws-eye .

# Run
./aws-eye                    # Default: ap-south-1
./aws-eye --region us-east-1 # Custom region
```

---

**You're all set!** 🚀

