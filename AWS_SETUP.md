# AWS Account Setup Guide for aws-eye CLI

This guide explains what you need from your AWS account to use the `aws-eye` CLI tool, including free tier requirements.

## ✅ What You Need

### 1. **AWS Account** (Free Tier is Sufficient)
- You can use a free tier AWS account
- No payment required for basic EC2 instance queries
- The `DescribeInstances` API call is **FREE** (it only reads metadata, doesn't create resources)

### 2. **IAM User with EC2 Read Permissions**
You need an IAM user with permission to describe EC2 instances.

### 3. **AWS Credentials**
Access Key ID and Secret Access Key for authentication.

---

## 🚀 Step-by-Step Setup

### Step 1: Create an IAM User (if you don't have one)

1. **Log in to AWS Console**
   - Go to https://console.aws.amazon.com
   - Sign in with your AWS account

2. **Navigate to IAM**
   - Search for "IAM" in the top search bar
   - Click on "IAM" service

3. **Create a New User**
   - Click "Users" in the left sidebar
   - Click "Create user"
   - Enter a username (e.g., `aws-eye-user`)
   - Click "Next"

4. **Attach Permissions**
   - Select "Attach policies directly"
   - Search for and select: **`AmazonEC2ReadOnlyAccess`**
     - This policy allows reading EC2 instance information (no charges)
   - Click "Next", then "Create user"

### Step 2: Create Access Keys

1. **Select Your User**
   - Click on the user you just created

2. **Create Access Key**
   - Go to the "Security credentials" tab
   - Scroll to "Access keys" section
   - Click "Create access key"
   - Select "Command Line Interface (CLI)" as the use case
   - Check the confirmation box and click "Next"
   - Optionally add a description tag, then click "Create access key"

3. **Save Your Credentials** ⚠️ **IMPORTANT**
   - **Access Key ID**: Copy and save this immediately
   - **Secret Access Key**: Copy and save this immediately
   - ⚠️ **You can only see the Secret Access Key once!** If you lose it, you'll need to create a new access key.

### Step 3: Configure AWS Credentials Locally

You have two options:

#### Option A: Environment Variables (Recommended for Testing)

```bash
export AWS_ACCESS_KEY_ID="your-access-key-id"
export AWS_SECRET_ACCESS_KEY="your-secret-access-key"
export AWS_DEFAULT_REGION="ap-south-1"  # Optional, can use --region flag
```

**For macOS/Linux:**
Add these to your `~/.zshrc` or `~/.bashrc`:
```bash
echo 'export AWS_ACCESS_KEY_ID="your-access-key-id"' >> ~/.zshrc
echo 'export AWS_SECRET_ACCESS_KEY="your-secret-access-key"' >> ~/.zshrc
source ~/.zshrc
```

#### Option B: AWS Credentials File (Recommended for Production)

Create/edit `~/.aws/credentials`:
```ini
[default]
aws_access_key_id = your-access-key-id
aws_secret_access_key = your-secret-access-key
```

Create/edit `~/.aws/config`:
```ini
[default]
region = ap-south-1
```

---

## 💰 Free Tier Information

### What's FREE:
- ✅ **EC2 DescribeInstances API calls** - Completely free (you're just reading metadata)
- ✅ **IAM users and policies** - Free
- ✅ **EC2 instance metadata queries** - Free

### What Costs Money:
- ❌ **Running EC2 instances** - Free tier includes 750 hours/month of t2.micro/t3.micro for 12 months
- ❌ **Stopped instances** - No charges (but they still exist and can be queried)

### Important Notes:
- **You can query instances even if you have none** - The API call will return an empty list
- **You can query instances in any region** - No charges for API calls
- **The CLI only READS data** - It never creates, modifies, or deletes anything

---

## 🧪 Testing Your Setup

1. **Verify credentials are set:**
   ```bash
   echo $AWS_ACCESS_KEY_ID
   echo $AWS_SECRET_ACCESS_KEY
   ```

2. **Run aws-eye:**
   ```bash
   ./aws-eye
   ```

3. **Expected Output:**
   - If you have instances: A table showing your EC2 instances
   - If you have no instances: "No instances found."
   - If credentials are wrong: An error message about authentication

---

## 🔒 Security Best Practices

1. **Never commit credentials to Git**
   - The `.gitignore` should exclude credential files
   - Use environment variables or AWS credentials file

2. **Use least privilege**
   - Only grant `AmazonEC2ReadOnlyAccess` (not full EC2 access)

3. **Rotate keys regularly**
   - Delete old access keys when creating new ones

4. **Use IAM roles for production**
   - For EC2 instances or Lambda, use IAM roles instead of access keys

---

## ❓ Troubleshooting

### Error: "failed to load AWS config"
- **Solution**: Check that your credentials are set correctly
- Verify environment variables or credentials file exists

### Error: "AccessDenied" or "UnauthorizedOperation"
- **Solution**: Ensure your IAM user has `AmazonEC2ReadOnlyAccess` policy attached

### Error: "InvalidClientTokenId"
- **Solution**: Your Access Key ID is incorrect

### Error: "SignatureDoesNotMatch"
- **Solution**: Your Secret Access Key is incorrect

### No instances shown
- **This is normal!** If you have no EC2 instances running, the list will be empty
- You can still use the CLI - it will just show "No instances found."

---

## 📚 Additional Resources

- [AWS Free Tier Details](https://aws.amazon.com/free/)
- [IAM User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users.html)
- [AWS CLI Configuration](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)

---

## ✅ Quick Checklist

- [ ] AWS account created (free tier OK)
- [ ] IAM user created
- [ ] `AmazonEC2ReadOnlyAccess` policy attached to user
- [ ] Access Key ID and Secret Access Key created and saved
- [ ] Credentials configured (environment variables or AWS credentials file)
- [ ] Tested `aws-eye` CLI tool

---

**You're all set!** The aws-eye CLI is ready to use with your free tier AWS account. 🎉

