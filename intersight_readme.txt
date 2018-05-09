# Import Powershell SDK module
Import-Module .\src\intersight\intersight.psd1

#Setup Handler for Certificate verification override
add-type @"
    using System.Net;
    using System.Security.Cryptography.X509Certificates;
    public class TrustAllCertsPolicy : ICertificatePolicy {
        public bool CheckValidationResult(
            ServicePoint srvPoint, X509Certificate certificate,
            WebRequest request, int certificateProblem) {
            return true;
        }
    }
"@
[System.Net.ServicePointManager]::CertificatePolicy = New-Object TrustAllCertsPolicy

#Steps to use the Powershell SDK
	# api key id value obtained from Intersight UI
$api_key_id="5a61b9896736327a31bbebff/5a61b7586736327a31bbeb74/5a7c3054647339736e57b82c"
	# location of .pem file saved for private key obtained from Intersight UI
$private_key_path="C:\\Users\\ratkv\\source\\repos\\key.pem"
	# Intersight URL
$intersightUrl = "https://ucs.cisco.com/api/v1"
New-IntersightApiClient $intersightUrl $private_key_path $api_key_id

# Sample script to add BIOS policies
for($i=0;$i -lt 5;$i++){
    $bios_policy =  New-BiosPolicy  -Name SampleBIOSpolicy$i
    Invoke-BiosPolicyApiBiosPoliciesPost $bios_policy
}