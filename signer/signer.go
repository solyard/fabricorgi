package signer

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fabricorgi/orgchecker"
)

func errorHandler(cmd *exec.Cmd) error {
	cmd.Stderr = os.Stderr
	_, err := cmd.Output()
	if err != nil {
		//log.Printf("%v", err)
		return err
	} else {
		//log.Printf("%v", string(out))
		return nil
	}
}

//SignAndAdd запуск ряда комманд для получения конфигурационного блока и внесения изменений
func SignAndAdd(payload *orgchecker.OrganizationConfig) error {

	// Инициализация переменной cmd
	var cmd *exec.Cmd
	var err error

	// Цикл по добавлению организации в канал
	cmd = exec.Command("bash", "-c", "peer channel fetch config config_block.pb --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -o ${FABRICORGI_ORDERER_IP} -c mainchannel")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", fmt.Sprintf(`jq -s '.[0] * {"channel_group":{"groups":{"Application":{"groups": {"'%s'":.[1]}}}}}' config.json ./organisation.json > modified_config.json`, payload.Policies.Readers.Policy.Value.Identities[0].Principal.MspIdentifier))
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input config.json --type common.Config --output config.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator compute_update --channel_id mainchannel --original config.pb --updated modified_config.pb --output organisation_update.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_decode --input organisation_update.pb --type common.ConfigUpdate | jq . > organisation_update.json")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", `echo '{"payload":{"header":{"channel_header":{"channel_id":"mainchannel", "type":2}},"data":{"config_update":'$(cat organisation_update.json)'}}}' | jq . > organisation_update_in_envelope.json`)
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input organisation_update_in_envelope.json --type common.Envelope --output organisation_update_in_envelope.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "peer channel signconfigtx --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -f organisation_update_in_envelope.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "peer channel update --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -f organisation_update_in_envelope.pb -c mainchannel -o ${FABRICORGI_ORDERER_IP}")
	err = errorHandler(cmd)

	if err != nil {
		return err
	}

	return nil
}

// SignAndRemove функция для удаления организации из канала
func SignAndRemove(payload *orgchecker.OrganizationRemove) error {

	// Инициализация переменной cmd
	var cmd *exec.Cmd
	var err error

	// Цикл по удалению организации из канала
	cmd = exec.Command("bash", "-c", "peer channel fetch config config_block.pb --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -o ${FABRICORGI_ORDERER_IP} -c mainchannel")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", fmt.Sprintf("jq 'del(.channel_group.groups.Application.groups.%s)' config.json > modified_config.json", payload.OrgName))
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input config.json --type common.Config --output config.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator compute_update --channel_id mainchannel --original config.pb --updated modified_config.pb --output update.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_decode --input update.pb --type common.ConfigUpdate | jq . > update.json")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", `echo '{"payload":{"header":{"channel_header":{"channel_id":"mainchannel", "type":2}},"data":{"config_update":'$(cat update.json)'}}}' | jq . > update_in_envelope.json`)
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "configtxlator proto_encode --input update_in_envelope.json --type common.Envelope --output update_in_envelope.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "peer channel signconfigtx --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -f update_in_envelope.pb")
	err = errorHandler(cmd)

	cmd = exec.Command("bash", "-c", "peer channel update --tls --cafile ${CORE_PEER_TLS_ROOTCERT_FILE} -f update_in_envelope.pb -c mainchannel -o ${FABRICORGI_ORDERER_IP}")
	err = errorHandler(cmd)

	if err != nil {
		return err
	}

	return nil
}
