package azurerm

import (
	"context"

	"github.com/sockeye44/terracognita/errcode"
	"github.com/sockeye44/terracognita/filter"
	"github.com/sockeye44/terracognita/provider"
	"github.com/pkg/errors"
)

func cacheVirtualNetworks(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]provider.Resource, error) {
	rs, err := a.cache.Get(rt)
	if err != nil {
		if errors.Cause(err) != errcode.ErrCacheKeyNotFound {
			return nil, errors.WithStack(err)
		}

		rs, err = virtualNetworks(ctx, a, rt, filters)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get virtual networks")
		}

		err = a.cache.Set(rt, rs)
		if err != nil {
			return nil, err
		}
	}

	return rs, nil
}
func getVirtualNetworkNames(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]string, error) {
	rs, err := cacheVirtualNetworks(ctx, a, rt, filters)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(rs))
	for _, i := range rs {
		names = append(names, i.Data().Get("name").(string))
	}

	return names, nil
}

func cacheVirtualMachines(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]provider.Resource, error) {
	rs, err := a.cache.Get(rt)
	if err != nil {
		if errors.Cause(err) != errcode.ErrCacheKeyNotFound {
			return nil, errors.WithStack(err)
		}

		rs, err = virtualMachines(ctx, a, rt, filters)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get virtual machines")
		}

		err = a.cache.Set(rt, rs)
		if err != nil {
			return nil, err
		}
	}

	return rs, nil
}
func getVirtualMachineNames(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]string, error) {
	rs, err := cacheVirtualMachines(ctx, a, rt, filters)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(rs))
	for _, i := range rs {
		names = append(names, i.Data().Get("name").(string))
	}

	return names, nil
}

func cacheWorkflows(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]provider.Resource, error) {
	rs, err := a.cache.Get(rt)
	if err != nil {
		if errors.Cause(err) != errcode.ErrCacheKeyNotFound {
			return nil, errors.WithStack(err)
		}

		rs, err = logicAppWorkflows(ctx, a, rt, filters)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get workflows")
		}

		err = a.cache.Set(rt, rs)
		if err != nil {
			return nil, err
		}
	}

	return rs, nil
}
func getWorkflowNames(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]string, error) {
	rs, err := cacheWorkflows(ctx, a, rt, filters)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(rs))
	for _, i := range rs {
		names = append(names, i.Data().Get("name").(string))
	}

	return names, nil
}

func cacheSecurityGroups(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]provider.Resource, error) {
	rs, err := a.cache.Get(rt)
	if err != nil {
		if errors.Cause(err) != errcode.ErrCacheKeyNotFound {
			return nil, errors.WithStack(err)
		}

		rs, err = networkSecurityGroups(ctx, a, rt, filters)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get Security Groups")
		}

		err = a.cache.Set(rt, rs)
		if err != nil {
			return nil, err
		}
	}

	return rs, nil
}
func getSecurityGroups(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]string, error) {
	rs, err := cacheSecurityGroups(ctx, a, rt, filters)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(rs))
	for _, i := range rs {
		names = append(names, i.Data().Get("name").(string))
	}

	return names, nil
}

func cacheRouteTables(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]provider.Resource, error) {
	rs, err := a.cache.Get(rt)
	if err != nil {
		if errors.Cause(err) != errcode.ErrCacheKeyNotFound {
			return nil, errors.WithStack(err)
		}

		rs, err = routeTables(ctx, a, rt, filters)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get routeTables")
		}

		err = a.cache.Set(rt, rs)
		if err != nil {
			return nil, err
		}
	}

	return rs, nil
}
func getRouteTables(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]string, error) {
	rs, err := cacheRouteTables(ctx, a, rt, filters)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(rs))
	for _, i := range rs {
		names = append(names, i.Data().Get("name").(string))
	}

	return names, nil
}

func cacheContainerRegistries(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]provider.Resource, error) {
	rs, err := a.cache.Get(rt)
	if err != nil {
		if errors.Cause(err) != errcode.ErrCacheKeyNotFound {
			return nil, errors.WithStack(err)
		}

		rs, err = containerRegistries(ctx, a, rt, filters)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get ContainerRegistries")
		}

		err = a.cache.Set(rt, rs)
		if err != nil {
			return nil, err
		}
	}

	return rs, nil
}
func getContainerRegistries(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]string, error) {
	rs, err := cacheContainerRegistries(ctx, a, rt, filters)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(rs))
	for _, i := range rs {
		names = append(names, i.Data().Get("name").(string))
	}

	return names, nil
}

func cacheStorageAccounts(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]provider.Resource, error) {
	rs, err := a.cache.Get(rt)
	if err != nil {
		if errors.Cause(err) != errcode.ErrCacheKeyNotFound {
			return nil, errors.WithStack(err)
		}

		rs, err = storageAccounts(ctx, a, rt, filters)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get storageAccounts")
		}

		err = a.cache.Set(rt, rs)
		if err != nil {
			return nil, err
		}
	}

	return rs, nil
}
func getStorageAccounts(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]string, error) {
	rs, err := cacheStorageAccounts(ctx, a, rt, filters)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(rs))
	for _, i := range rs {
		names = append(names, i.Data().Get("name").(string))
	}

	return names, nil
}

func cacheMariadbServers(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]provider.Resource, error) {
	rs, err := a.cache.Get(rt)
	if err != nil {
		if errors.Cause(err) != errcode.ErrCacheKeyNotFound {
			return nil, errors.WithStack(err)
		}

		rs, err = mariadbServers(ctx, a, rt, filters)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get MariaDB Servers")
		}

		err = a.cache.Set(rt, rs)
		if err != nil {
			return nil, err
		}
	}

	return rs, nil
}
func getMariadbServers(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]string, error) {
	rs, err := cacheMariadbServers(ctx, a, rt, filters)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(rs))
	for _, i := range rs {
		names = append(names, i.Data().Get("name").(string))
	}

	return names, nil
}

func cacheMysqlServers(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]provider.Resource, error) {
	rs, err := a.cache.Get(rt)
	if err != nil {
		if errors.Cause(err) != errcode.ErrCacheKeyNotFound {
			return nil, errors.WithStack(err)
		}

		rs, err = mysqlServers(ctx, a, rt, filters)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get MySQL Servers")
		}

		err = a.cache.Set(rt, rs)
		if err != nil {
			return nil, err
		}
	}

	return rs, nil
}
func getMysqlServers(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]string, error) {
	rs, err := cacheMysqlServers(ctx, a, rt, filters)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(rs))
	for _, i := range rs {
		names = append(names, i.Data().Get("name").(string))
	}

	return names, nil
}

func cachePostgresqlServers(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]provider.Resource, error) {
	rs, err := a.cache.Get(rt)
	if err != nil {
		if errors.Cause(err) != errcode.ErrCacheKeyNotFound {
			return nil, errors.WithStack(err)
		}

		rs, err = postgresqlServers(ctx, a, rt, filters)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get PostgreSQL Servers")
		}

		err = a.cache.Set(rt, rs)
		if err != nil {
			return nil, err
		}
	}

	return rs, nil
}
func getPostgresqlServers(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]string, error) {
	rs, err := cachePostgresqlServers(ctx, a, rt, filters)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(rs))
	for _, i := range rs {
		names = append(names, i.Data().Get("name").(string))
	}

	return names, nil
}

func cacheSQLServers(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]provider.Resource, error) {
	rs, err := a.cache.Get(rt)
	if err != nil {
		if errors.Cause(err) != errcode.ErrCacheKeyNotFound {
			return nil, errors.WithStack(err)
		}

		rs, err = sqlServers(ctx, a, rt, filters)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get SQL Servers")
		}

		err = a.cache.Set(rt, rs)
		if err != nil {
			return nil, err
		}
	}

	return rs, nil
}
func getSQLServers(ctx context.Context, a *azurerm, rt string, filters *filter.Filter) ([]string, error) {
	rs, err := cacheSQLServers(ctx, a, rt, filters)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(rs))
	for _, i := range rs {
		names = append(names, i.Data().Get("name").(string))
	}

	return names, nil
}
