$version = "0.18.1"
if (-Not $env:VERSION -eq $null) {
  $version = "$env:VERSION"
}
$url = "https://github.com/tinned-fish/gofish/releases/download/v${version}/gofish-${version}-windows-amd64.tar.gz"

if ($env:TEMP -eq $null) {
  $env:TEMP = Join-Path $env:SystemDrive 'temp'
}
$tempDir = Join-Path $env:TEMP 'Fish'
if (![System.IO.Directory]::Exists($tempDir)) {[void][System.IO.Directory]::CreateDirectory($tempDir)}
$file = Join-Path $env:TEMP "gofish-${version}-windows-amd64.tar.gz"

# Download gofish
Write-Output "Downloading '$url'"
(new-object System.Net.WebClient).DownloadFile($url, $file)

$installPath = "$env:SYSTEMDRIVE\ProgramData\bin"
if (![System.IO.Directory]::Exists($installPath)) {[void][System.IO.Directory]::CreateDirectory($installPath)}
Write-Output "Preparing to install into '$installPath'"

$binaryPath = "$installPath\gofish.exe"
Expand-Archive -Path "$file" -DestinationPath "$tempDir" -Force
if ([System.IO.File]::Exists("$binaryPath")) {[void][System.IO.File]::Delete("$binaryPath")}
Move-Item -Path "$tempDir\windows-amd64\gofish.exe" -Destination "$binaryPath"

# Add gofish to the path
if ($($env:Path).ToLower().Contains($($installPath).ToLower()) -eq $false) {
  Write-Output "Adding '$installPath' to system PATH"
  $newPath = [Environment]::GetEnvironmentVariable('Path',[System.EnvironmentVariableTarget]::Machine) + ";$installPath";
  [Environment]::SetEnvironmentVariable('Path',$newPath,[System.EnvironmentVariableTarget]::Machine);
}

Write-Output "gofish installed to '$binaryPath'"
Write-Output "Restart your terminal, then run 'gofish init' to get started!"
