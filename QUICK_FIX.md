# Quick Fix: UnauthorizedOperation Error

## Your Current Error

```
User: arn:aws:iam::225989378089:user/neeraj115 is not authorized to perform: ec2:DescribeInstances
```

**Problem**: Your IAM user `neeraj115` doesn't have permission to read EC2 instances.

---

## ✅ Solution: Add EC2 Read Permissions

### Option 1: Using AWS Console (Easiest)

1. **Go to AWS Console**
   - Visit: https://console.aws.amazon.com
   - Sign in with your AWS account

2. **Navigate to IAM**
   - Search for "IAM" in the top search bar
   - Click on "IAM" service

3. **Find Your User**
   - Click "Users" in the left sidebar
   - Search for and click on: **`neeraj115`**

4. **Add Permissions**
   - Click the "Add permissions" button
   - Select "Attach policies directly"
   - In the search box, type: `AmazonEC2ReadOnlyAccess`
   - Check the box next to **`AmazonEC2ReadOnlyAccess`**
   - Click "Next", then "Add permissions"

5. **Verify**
   - You should see the policy attached under "Permissions" tab
   - The policy should show: `AmazonEC2ReadOnlyAccess`

6. **Test Again**
   ```bash
   ./aws-eye
   ```

---

### Option 2: Using AWS CLI (If you have admin access)

If you have AWS CLI installed and admin permissions:

```bash
aws iam attach-user-policy \
    --user-name neeraj115 \
    --policy-arn arn:aws:iam::aws:policy/AmazonEC2ReadOnlyAccess
```

Then test:
```bash
./aws-eye
```

---

## ⚠️ If You Don't Have Permission to Modify IAM

If you're not the account owner or don't have IAM admin permissions:

1. **Contact your AWS account administrator**
   - Ask them to attach `AmazonEC2ReadOnlyAccess` policy to user `neeraj115`
   - Or ask them to create a new IAM user with EC2 read permissions

2. **Alternative: Use a Different User**
   - If you have access to another IAM user with EC2 permissions
   - Update your `~/.aws/credentials` file with that user's keys:
     ```ini
     [default]
     aws_access_key_id = NEW_ACCESS_KEY_ID
     aws_secret_access_key = NEW_SECRET_ACCESS_KEY
     ```

---

## Verify the Fix

After adding permissions, wait 1-2 minutes for changes to propagate, then:

```bash
./aws-eye
```

**Expected Results:**
- ✅ **Success**: Shows "Fetching EC2 instances..." and either a table or "No instances found."
- ❌ **Still Error**: Wait a bit longer (up to 5 minutes) for IAM changes to propagate globally

---

## What This Policy Allows

The `AmazonEC2ReadOnlyAccess` policy allows:
- ✅ Reading EC2 instance information (DescribeInstances)
- ✅ Viewing instance states, IPs, availability zones
- ❌ **NO** ability to create, modify, or delete instances
- ❌ **NO** charges - it's just reading metadata

**This is safe and free to use!**

---

## Check Current Permissions

To see what permissions your user currently has:

1. Go to IAM Console → Users → `neeraj115`
2. Click "Permissions" tab
3. Review the policies listed

If you see `AmazonEC2ReadOnlyAccess` or any policy with `ec2:DescribeInstances` permission, you're good to go!

